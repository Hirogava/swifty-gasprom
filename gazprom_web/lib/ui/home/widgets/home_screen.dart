import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:gazprom_web/domain/models/gift.dart';
import 'package:gazprom_web/ui/core/themes/theme.dart';
import 'package:gazprom_web/ui/core/ui/universal_container.dart';

import '../../core/themes/colors.dart';
import '../../core/ui/slanted_app_bar.dart';
import '../../core/ui/svg_picture_widget.dart';
import 'card_widget.dart';
import 'character.dart';
import 'gift_widget.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.transparent,
      appBar: const SlantedAppBar(title: 'НОВОСТИ'),
      body: Stack(
        children: [
          SingleChildScrollView(
            child: Padding(
              padding: EdgeInsets.symmetric(horizontal: 16, vertical: 40),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  AnimatedSmiley(),
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
                    children: [
                      Expanded(
                        child: UniversalContainer(
                          borderRadius: 15,
                          children: [
                            Center(child: Text('Активы', style: AppTextStyles.style3)),
                          ],
                        ),
                      ),
                      Icon(Icons.chevron_right),
                      Expanded(
                        child: UniversalContainer(
                          borderRadius: 15,
                          children: [
                            Center(child: Text('Пассивы', style: AppTextStyles.style3)),
                          ],
                        ),
                      ),
                    ],
                  ),
                  SizedBox(height: 20),
                  UniversalContainer(borderRadius: 20, children: [Loan()]),
                ],
              ),
            ),
          ),
          Positioned(
            top: 50,
            left: 20,
            child: SvgPicture.asset(
              'assets/nav_icons/login.svg',
              width: 35,
              height: 35,
              colorFilter: ColorFilter.mode(AppColors.white, BlendMode.srcIn),
            ),
          ),
        ],
      ),
    );
  }
}

class Loan extends StatelessWidget {
  const Loan({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          children: [
            SvgPictureWidget(path: 'assets/other/card_logo.svg', width: 216, height: 31, color: AppColors.white),
            Spacer(),
            Column(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Text('0 руб', style: AppTextStyles.style4),
                Text('Платеж в год', style: AppTextStyles.style5),
              ],
            ),
          ],
        ),
        Text('Button', style: AppTextStyles.style3),
        SizedBox(height: 10),
        Text('Button', style: AppTextStyles.style3),
        UniversalContainer(
          borderRadius: 100,
          children: [Text('Оформить кредит', style: AppTextStyles.style4)],
        ),
      ],
    );
  }
}
