import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/core/themes/colors.dart';

import '../../../domain/models/gift.dart';

class GiftWidget extends StatelessWidget {
  final Gift gift;
  const GiftWidget({super.key, required this.gift});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 5, horizontal: 10),
      decoration: BoxDecoration(
        color: AppColors.white,
        borderRadius: BorderRadius.circular(15),
      ),
      child: Column(
        children: [
          Text(gift.title),
          SizedBox(height: 10),
          Text(gift.description),
          SizedBox(height: 10),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              SizedBox(width: 10),
              Expanded(
                child: ButtonWidget(
                  title: 'Отказать',
                  gradient: AppColors.purpleGradient,
                ),
              ),
              SizedBox(width: 10),
              Expanded(child: ButtonWidget(title: 'Согласиться')),
              SizedBox(width: 10),
            ],
          ),
        ],
      ),
    );
  }
}

class ButtonWidget extends StatelessWidget {
  final String title;
  final Gradient? gradient;
  const ButtonWidget({super.key, required this.title, this.gradient});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {},
      child: Container(
        padding: EdgeInsets.symmetric(horizontal: 5, vertical: 10),
        decoration: BoxDecoration(
          color: gradient == null ? AppColors.blue1 : null,
          gradient: gradient,
          borderRadius: BorderRadius.circular(15),
        ),
        alignment: Alignment.center,
        child: Text(
          title,
          style: const TextStyle(color: AppColors.white),
          textAlign: TextAlign.center,
        ),
      ),
    );
  }
}
