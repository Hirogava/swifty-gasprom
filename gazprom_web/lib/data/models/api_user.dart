import '../../domain/models/user.dart';

class ApiUser {
  final String id;
  final String bankUserId;
  final String accessToken;
  final String refreshToken;

  ApiUser({
    required this.id,
    required this.bankUserId,
    required this.accessToken,
    required this.refreshToken,
  });

  factory ApiUser.fromJson(Map<String, dynamic> json) {
    return ApiUser(
      id: json['id'],
      bankUserId: json['bank_user_id'],
      accessToken: json['token']['access_token'],
      refreshToken: json['token']['refresh_token'],
    );
  }

  User toDomain() {
    return User(
      id: id,
      bankUserId: bankUserId,
      accessToken: accessToken,
      refreshToken: refreshToken,
      position: null,
      happiness: 0,
    );
  }
}
