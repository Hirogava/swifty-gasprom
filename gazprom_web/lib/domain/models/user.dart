import 'position.dart';

class User {
  final String id;
  final String bankUserId;
  final String accessToken;
  final String refreshToken;
  final Position? position;
  final int happiness;

  User({required this.id, required this.bankUserId, required this.accessToken, required this.refreshToken, required this.position, required this.happiness});
}