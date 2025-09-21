import 'package:flutter/material.dart';

import '../../core/themes/colors.dart';

class RiskScreen extends StatelessWidget {
  const RiskScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Background(
      child: Scaffold(
        backgroundColor: AppColors.transparent,
        body: Center(
          child: Text('Risk Screen', style: TextStyle(color: AppColors.white),),
        ),
      ),
    );
  }
}