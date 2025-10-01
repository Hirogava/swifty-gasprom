import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/core/ui/button_widget.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../core/themes/colors.dart';
import '../../core/themes/theme.dart';

class WelcomeScreen extends StatelessWidget {
  const WelcomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Background(
      child: Scaffold(
        backgroundColor: AppColors.transparent,
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: 16, vertical: 100),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text(
                'GOTOCASH',
                style: GoogleFonts.righteous(
                  fontSize: 24,
                  color: AppColors.white,
                  fontWeight: AppFontWeights.bolt,
                ),
              ),
              Text(
                'by',
                style: GoogleFonts.righteous(
                  fontSize: 18,
                  color: AppColors.white,
                  fontWeight: AppFontWeights.bolt300,
                ),
              ),
              Text(
                'SWIFTLY',
                style: GoogleFonts.righteous(
                  fontSize: 20,
                  color: AppColors.white,
                  fontWeight: AppFontWeights.bolt300,
                ),
              ),
              SizedBox(height: 50),
              Text(
                'Добро пожаловать в игру по финансовой грамотности',
                style: GoogleFonts.righteous(
                  fontSize: 16,
                  color: AppColors.white,
                  fontWeight: AppFontWeights.normal,
                ),
              ),
              SizedBox(height: 20),
              Text(
                'Благодаря ГАЗПРОМБАНК ТЕХ у вас есть возможность погрузиться в захватывающий мир финансов и научитесь управлять своими деньгами, инвестировать и планировать бюджет с помощью интерактивной и увлекательной игры. Улучшайте навыки, получайте награды и достигайте новых уровней!',
                style: GoogleFonts.inter(
                  fontSize: 12,
                  color: AppColors.white,
                  fontWeight: AppFontWeights.normal,
                ),
              ),
              SizedBox(height: 10),
              Text(
                'Игра разработана специально для пользователей ГАЗПРОМБАНК ТЕХ',
                style: GoogleFonts.inter(
                  fontSize: 12,
                  color: AppColors.white,
                  fontWeight: AppFontWeights.normal,
                ),
              ),
              Spacer(),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  ButtonWidget(
                    buttonText: 'Назад',
                    textColor: AppColors.white,
                    borderColor: AppColors.white,
                    gradient: AppColors.blackGradient3,
                    width: 138,
                    height: 45,
                    pathScreen: '/',
                  ),
                  ButtonWidget(
                    buttonText: 'Далее',
                    textColor: AppColors.white,
                    borderColor: AppColors.white,
                    gradient: AppColors.blackGradient3,
                    width: 138,
                    height: 45,
                    pathScreen: '/login',
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
