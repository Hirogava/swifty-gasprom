import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:gazprom_web/providers/current_user_provider.dart';
import 'package:go_router/go_router.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../providers/user_repo_provider.dart';
import '../../core/themes/colors.dart';
import '../../core/themes/theme.dart';
import '../../core/ui/button_widget.dart';

import 'dart:developer' as dev;

class LoginScreen extends ConsumerWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    Future<void> login() async {
      try {
        final repo = ref.read(userRepoProvider);
        // тут можешь сгенерить рандомный id
        final user = await repo.login('123');
        if (!context.mounted) return;
        ref.read(currentUserProvider.notifier).state = user;

        dev.log("user id = ${user.id}");

        // навигация после успешного логина
        if (context.mounted) {
          context.go('/home');
        }
      } catch (e, s) {
        dev.log("Login failed", error: e, stackTrace: s);
        ScaffoldMessenger.of(
          context,
        ).showSnackBar(SnackBar(content: Text("Ошибка авторизации")));
      }
    }

    return Background(
      child: Scaffold(
        backgroundColor: AppColors.transparent,
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: 16, vertical: 100),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text(
                'Авторизоваться через банк',
                style: GoogleFonts.inter(
                  fontSize: 20,
                  color: AppColors.white,
                  fontWeight: AppFontWeights.normal,
                ),
              ),
              SizedBox(height: 20),
              ButtonWidget(
                buttonText: 'АВТОРИЗОВАТЬСЯ',
                textColor: AppColors.white,
                borderColor: AppColors.white,
                gradient: AppColors.blackGradient3,
                width: 333,
                height: 45,
                pathScreen: '/home',
                hm: login,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
