import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:gazprom_web/domain/models/gift.dart';
import 'package:gazprom_web/ui/core/themes/theme.dart';
import 'package:gazprom_web/ui/core/ui/universal_container.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../core/themes/colors.dart';
import '../../core/ui/svg_picture_widget.dart';
import 'card_widget.dart';
import 'gift_widget.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.transparent,
      appBar: AppBar(
        title: Stack(
          children: [
            Align(
              alignment: Alignment.centerLeft,
              child: SvgPicture.asset(
                'assets/nav_icons/login.svg',
                width: 35,
                height: 35,
                colorFilter: ColorFilter.mode(AppColors.white, BlendMode.srcIn),
              ),
            ),
            Align(
              alignment: Alignment.center,
              child: Text(
                'Главная',
                style: GoogleFonts.roboto(fontSize: 24, color: AppColors.white),
              ),
            ),
          ],
        ),
        backgroundColor: AppColors.transparent,
        shape: const Border(bottom: BorderSide(color: Colors.white)),
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 16, vertical: 40),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              SizedBox(height: 20),
              CardWidget(balance: '300 000'),
              SizedBox(height: 20),
              GiftWidget(
                gift: Gift(
                  title: 'Внезапный подарок родственника.',
                  description:
                      'День шел как обычно, но тут вам набрала ваша тетя. Так как все ее деньги находятся в обороте, она просит вас одолжить ей 250 000 рублей на операцию.',
                ),
              ),
              SizedBox(height: 20),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  UniversalContainer(
                    gradient: AppColors.blackGradient3,
                    width: 178,
                    borderRadius: 15,
                    children: [
                      Center(
                        child: Text('Активы', style: AppTextStyles.style3),
                      ),
                    ],
                  ),
                  Icon(Icons.chevron_right, color: AppColors.white, size: 50),
                  UniversalContainer(
                    gradient: AppColors.blackGradient3,
                    width: 178,
                    borderRadius: 15,
                    children: [
                      Center(
                        child: Text('Пассивы', style: AppTextStyles.style3),
                      ),
                    ],
                  ),
                ],
              ),
              SizedBox(height: 20),
              Center(
                child: UniversalContainer(gradient: AppColors.blackGradient3, borderRadius: 20, children: [Loan()]),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class Loan extends StatefulWidget {
  const Loan({super.key});

  @override
  State<Loan> createState() => _LoanState();
}

class _LoanState extends State<Loan> {
  double _credit = 100000;
  double _term = 12;
  final double _annualRate = 0.12;

  double get monthlyPayment {
    final i = _annualRate / 12;
    final n = _term.toInt();

    if (i == 0) return _credit / n;

    return _credit * i / (1 - (1 / (pow(1 + i, n))));
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          children: [
            SvgPictureWidget(
              path: 'assets/other/card_logo.svg',
              width: 216,
              height: 31,
              color: AppColors.white,
            ),
            const Spacer(),
            Column(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Text(
                  '${monthlyPayment.toStringAsFixed(0)} руб/мес',
                  style: GoogleFonts.robotoMono(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                    color: AppColors.white,
                  ),
                ),
                Text('Платеж в месяц', style: AppTextStyles.style5),
              ],
            ),
          ],
        ),
        const SizedBox(height: 10),

        Text('Сумма кредита', style: AppTextStyles.style8),
        Text(
          '${_credit.toInt()} руб',
          style: GoogleFonts.robotoMono(
            fontSize: 13,
            fontWeight: AppFontWeights.bolt500,
            color: AppColors.white,
          ),
        ),
        SliderTheme(
          data: SliderTheme.of(
            context,
          ).copyWith(showValueIndicator: ShowValueIndicator.never),
          child: Slider(
            activeColor: AppColors.white,
            value: _credit,
            min: 10000,
            max: 5000000,
            divisions: ((5000000 - 10000) ~/ 1000),
            label: '${_credit.toInt()}',
            onChanged: (value) => setState(() => _credit = value),
          ),
        ),

        const SizedBox(height: 10),

        Text('Срок (в месяцах)', style: AppTextStyles.style8),
        Text(
          '${_term.toInt()} мес.',
          style: GoogleFonts.robotoMono(
            fontSize: 13,
            fontWeight: AppFontWeights.bolt500,
            color: AppColors.white,
          ),
        ),
        SliderTheme(
          data: SliderTheme.of(
            context,
          ).copyWith(showValueIndicator: ShowValueIndicator.never),
          child: Slider(
            activeColor: AppColors.white,
            value: _term,
            min: 12,
            max: 120,
            divisions: (120 - 12),
            label: '${_term.toInt()} мес.',
            onChanged: (value) => setState(() => _term = value),
          ),
        ),

        const SizedBox(height: 10),

        Center(
          child: UniversalContainer(
            gradient: AppColors.blackGradient3,
            borderRadius: 100,
            isButton: true,
            onTap: () {
              debugPrint(
                'Оформляем кредит $_credit руб на $_term мес. '
                'Платеж = ${monthlyPayment.toStringAsFixed(2)} руб/мес',
              );
            },
            children: [Text('Оформить кредит', style: AppTextStyles.style4)],
          ),
        ),
      ],
    );
  }
}
