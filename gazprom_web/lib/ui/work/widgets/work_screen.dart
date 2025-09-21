import 'package:flutter/material.dart';

import '../../core/themes/colors.dart';

class WorkScreen extends StatelessWidget {
  const WorkScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Background(
      child: Scaffold(
        backgroundColor: AppColors.transparent,
        body: Center(
          child: Text('Work Screen', style: TextStyle(color: AppColors.white),),
        ),
      ),
    );
  }
}