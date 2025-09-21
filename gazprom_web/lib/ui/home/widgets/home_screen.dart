import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:gazprom_web/domain/models/gift.dart';

import '../../core/themes/colors.dart';
import 'card_widget.dart';
import 'character.dart';
import 'gift_widget.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.blue2,
      body: Stack(
        children: [
          Padding(
            padding: EdgeInsets.symmetric(horizontal: 16, vertical: 40),
            child: SingleChildScrollView(
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
                  GiftWidget(
                    gift: Gift(
                      title: 'Внезапный подарок родственника.',
                      description:
                          'День шел как обычно, но тут вам набрала ваша тетя. Так как все ее деньги находятся в обороте, она просит вас одолжить ей 250 000 рублей на операцию.',
                    ),
                  ),
                ],
              ),
            ),
          ),
          Positioned(
            top: 50,
            left: 20,
            child: SvgPicture.asset(
              'assets/icons/login.svg',
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
