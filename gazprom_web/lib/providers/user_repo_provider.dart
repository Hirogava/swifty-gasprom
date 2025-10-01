import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../data/repositories/user_repo.dart';
import 'api_service_provider.dart';

final userRepoProvider = Provider<UserRepo>((ref) {
  return UserRepo(ref.read(apiServiceProvider),);
});