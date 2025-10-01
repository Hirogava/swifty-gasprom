import 'package:flutter/material.dart';

import '../../core/themes/colors.dart';
import '../../core/themes/theme.dart';
import '../../core/ui/svg_picture_widget.dart';

class CardWidget extends StatelessWidget {
  final String balance;
  const CardWidget({super.key, required this.balance});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(horizontal: 15, vertical: 20),
      width: 325,
      height: 210,
      decoration: BoxDecoration(
        gradient: AppColors.blackGradient1,
        borderRadius: BorderRadius.circular(15),
        boxShadow: [
          BoxShadow(
            color: AppColors.black20,
            spreadRadius: 0,
            blurRadius: 8,
            offset: const Offset(0, -5),
          ),
        ],
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SvgPictureWidget(path: 'assets/other/card_logo.svg', width: 216, height: 31, color: AppColors.white),
          SizedBox(height: 20),
          Text('БАЛАНС', style: AppTextStyles.style1),
          Text('$balance РУБ', style: AppTextStyles.style1),
          Spacer(),
          Center(child: Text('SUPREME', style: AppTextStyles.style2)),
        ],
      ),
    );
  }
}
