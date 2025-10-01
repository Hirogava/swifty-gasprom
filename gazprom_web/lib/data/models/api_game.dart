import '../../domain/models/game.dart';

class ApiGame {
  final int id;
  final String userId;
  final double money;
  final int happiness;
  final int month;
  final String status;
  final int naturalExpenses;

  ApiGame({
    required this.id,
    required this.userId,
    required this.money,
    required this.happiness,
    required this.month,
    required this.status,
    required this.naturalExpenses,
  });

  factory ApiGame.fromJson(Map<String, dynamic> json) {
    return ApiGame(
      id: json['id'],
      userId: json['user_id'],
      money: (json['money'] as num).toDouble(),
      happiness: json['happiness'],
      month: json['month'],
      status: json['status'],
      naturalExpenses: json['natural_expenses'],
    );
  }

  Game toDomain() {
    return Game(
      id: id,
      userId: userId,
      money: money,
      happiness: happiness,
      month: month,
      status: status,
      naturalExpenses: naturalExpenses,
    );
  }
}