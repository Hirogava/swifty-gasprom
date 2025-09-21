import 'package:flutter/material.dart';

import '../../core/themes/colors.dart';

class NewsScreen extends StatelessWidget {
  const NewsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Background(
      child: Scaffold(
        backgroundColor: AppColors.transparent,
        body: Center(
          child: Text('News Screen', style: TextStyle(color: AppColors.white),),
        ),
      ),
    );
  }
}