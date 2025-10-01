class AppException implements Exception {
  final String message;
  final String? code;
  
  const AppException(this.message, {this.code});

  @override
  String toString() => code != null 
      ? '$runtimeType($code): $message' 
      : '$runtimeType: $message';
}

// Сетевые ошибки
class NetworkException extends AppException {
  const NetworkException(super.message, {super.code});
}

class NoInternetException extends NetworkException {
  const NoInternetException() 
      : super('No internet connection', code: 'NO_INTERNET');
}

class ApiTimeoutException extends NetworkException {
  const ApiTimeoutException() 
      : super('Request timeout', code: 'TIMEOUT');
}

// Ошибки API
class ApiException extends AppException {
  final int? statusCode;
  
  const ApiException(
    super.message, {
    this.statusCode,
    super.code = 'API_ERROR',
  });
}

// Ошибки данных
class DataValidationException extends AppException {
  final String? field;
  
  const DataValidationException(
    super.message, {
    this.field,
    super.code = 'VALIDATION_ERROR',
  });
  
  @override
  String toString() => field != null
      ? '${super.toString()} (field: $field)'
      : super.toString();
}