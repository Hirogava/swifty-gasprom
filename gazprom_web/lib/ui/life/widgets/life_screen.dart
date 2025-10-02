import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/core/ui/universal_container.dart';
import 'package:gazprom_web/ui/home/widgets/card_widget.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../core/themes/colors.dart';

class LifeScreen extends StatelessWidget {
  const LifeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Background(
      child: Scaffold(
        appBar: AppBar(
          title: Text(
            'Жизнь',
            style: GoogleFonts.roboto(fontSize: 24, color: AppColors.white),
          ),
          backgroundColor: AppColors.transparent,
          shape: const Border(bottom: BorderSide(color: Colors.white)),
        ),
        backgroundColor: AppColors.transparent,
        body: SingleChildScrollView(
          child: Padding(
            padding: EdgeInsets.symmetric(horizontal: 16, vertical: 40),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                CardWidget(balance: '100'),
                SizedBox(height: 10),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    Text('sphera'),
                    Text('happinies'),
                  ],
                ),
                SizedBox(height: 10),
                UniversalContainer(borderRadius: 15, width: double.infinity, gradient: AppColors.blackGradient3, children: [Text('Развлечения и отдых', style: GoogleFonts.inter(fontSize: 24, color: AppColors.white),)]),
                SizedBox(height: 10),
                UniversalContainer(borderRadius: 15, width: double.infinity, gradient: AppColors.blackGradient3, children: [Text('Техника и гаджеты', style: GoogleFonts.inter(fontSize: 24, color: AppColors.white),)]),
                SizedBox(height: 10),
                UniversalContainer(borderRadius: 15, width: double.infinity, gradient: AppColors.blackGradient3, children: [Text('Подарки и соц.', style: GoogleFonts.inter(fontSize: 24, color: AppColors.white),)]),
                SizedBox(height: 10),
                UniversalContainer(borderRadius: 15, width: double.infinity, gradient: AppColors.blackGradient3, children: [Text('Образование и саморазвитие', style: GoogleFonts.inter(fontSize: 24, color: AppColors.white),)]),
                SizedBox(height: 10),
                UniversalContainer(borderRadius: 15, width: double.infinity, gradient: AppColors.blackGradient3, children: [Text('Здоровье и уход', style: GoogleFonts.inter(fontSize: 24, color: AppColors.white),)]),
                SizedBox(height: 10),
                UniversalContainer(borderRadius: 15, width: double.infinity, gradient: AppColors.blackGradient3, children: [Text('Повседневные радости', style: GoogleFonts.inter(fontSize: 24, color: AppColors.white),)]),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
