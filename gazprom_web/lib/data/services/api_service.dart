import 'package:dio/dio.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:gazprom_web/core/exceptions.dart';

import 'dart:developer' as dev;

import '../models/api_game.dart';
import '../models/api_user.dart';

class ApiService {
  final Dio dio;
  final InternetConnectionChecker internetConnectionChecker;

  ApiService({Dio? dio, InternetConnectionChecker? internetConnectionChecker})
    : dio = dio ?? Dio(BaseOptions(baseUrl: 'https://api.web24.team/api/v1/')),
      internetConnectionChecker =
          internetConnectionChecker ?? InternetConnectionChecker.instance;

  Future<ApiUser> login(String id) async {
    final bool isConnected = await internetConnectionChecker.hasConnection;

    if (!isConnected) {
      dev.log('No internet connection');
      throw const NoInternetException();
    }

    try {
      final response = await dio.post('login', data: {'bank_user_id': id});

      final data = response.data['user'];

      if (response.statusCode == 200) {
        dev.log('Успешно: ${response.data}');
        return ApiUser.fromJson(data);
      } else {
        throw ApiException(
          'Invalid status code',
          statusCode: response.statusCode,
        );
      }
    } on DioException catch (e, s) {
      dev.log('Dio error', error: e, stackTrace: s);

      if (e.type == DioExceptionType.connectionTimeout ||
          e.type == DioExceptionType.receiveTimeout) {
        throw const ApiTimeoutException();
      }

      if (e.response != null) {
        throw ApiException(
          'Server error: ${e.response?.statusMessage ?? 'Unknown'}',
          statusCode: e.response?.statusCode,
        );
      }

      throw NetworkException('Network error: ${e.message}');
    } catch (e, s) {
      dev.log('Unexpected API error', error: e, stackTrace: s);
      throw const NetworkException('Unexpected network error');
    }
  }

  Future<String> logout(String accessToken) async {
    try {
      final response = await dio.post('logout', data: {'logout': accessToken});

      final data = response.data['user'];

      if (response.statusCode == 200) {
        dev.log('Успешно: ${response.data}');
        return data;
      } else {
        throw ApiException(
          'Invalid status code',
          statusCode: response.statusCode,
        );
      }
    } on DioException catch (e, s) {
      dev.log('Dio error', error: e, stackTrace: s);

      if (e.type == DioExceptionType.connectionTimeout ||
          e.type == DioExceptionType.receiveTimeout) {
        throw const ApiTimeoutException();
      }

      if (e.response != null) {
        throw ApiException(
          'Server error: ${e.response?.statusMessage ?? 'Unknown'}',
          statusCode: e.response?.statusCode,
        );
      }

      throw NetworkException('Network error: ${e.message}');
    } catch (e, s) {
      dev.log('Unexpected API error', error: e, stackTrace: s);
      throw const NetworkException('Unexpected network error');
    }
  }

  Future<ApiGame> startGame() async {
    try {
      final response = await dio.get('start');

      _validateResponse(response);

      final data = response.data['game'];

      return ApiGame.fromJson(data);
    } on DioException catch (e) {
      throw Exception('Failed to start game: ${e.message}');
    }
  }

  void _validateResponse(Response response) {
    if (response.statusCode != 200) {
      throw ApiException(
        'Invalid status code',
        statusCode: response.statusCode,
      );
    }

    if (response.data == null || response.data['game'] == null) {
      throw const DataValidationException('Invalid response format');
    }
  }
}
