import 'package:flutter/material.dart';

import '../../core/themes/colors.dart';

class LifeScreen extends StatelessWidget {
  const LifeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Background(
      child: Scaffold(
        backgroundColor: AppColors.transparent,
        body: Center(
          child: Text('Life Screen', style: TextStyle(color: AppColors.white),),
        ),
      ),
    );
  }
}