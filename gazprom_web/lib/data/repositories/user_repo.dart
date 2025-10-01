import '../../core/exceptions.dart';
import '../../domain/models/user.dart';
import '../services/api_service.dart';

import 'dart:developer' as dev;

class UserRepo {
  final ApiService apiService;

  UserRepo(this.apiService);

  Future<User> login(String bankUserId) async {
    try {
      final apiUser = await apiService.login(bankUserId);
      return apiUser.toDomain();
    } on AppException {
      rethrow;
    } catch (e, s) {
      dev.log('Unexpected repository error', error: e, stackTrace: s);
      throw const AppException('Failed to login');
    }
  }
}
