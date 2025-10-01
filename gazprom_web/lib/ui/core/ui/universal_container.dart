import 'package:flutter/material.dart';

import '../themes/colors.dart';

class UniversalContainer extends StatelessWidget {
  final double? width;
  final double? height;
  final double borderRadius;
  final bool isButton;
  final VoidCallback? onTap;
  final LinearGradient? gradient;
  final List<Widget> children;

  const UniversalContainer({
    super.key,
    this.width,
    this.height,
    required this.borderRadius,
    this.isButton = false,
    this.onTap,
    this.gradient,
    required this.children,
  });

  @override
  Widget build(BuildContext context) {
    final child = Container(
      padding: const EdgeInsets.all(10),
      width: width,
      height: height,
      decoration: BoxDecoration(
        color: gradient == null ? AppColors.transparent : null,
        gradient: gradient,
        borderRadius: BorderRadius.circular(borderRadius),
        border: Border.all(color: AppColors.white),
      ),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          for (int i = 0; i < children.length; i++) ...[
            children[i],
            if (i < children.length - 1) const SizedBox(height: 10),
          ],
        ],
      ),
    );

    return isButton ? GestureDetector(onTap: onTap, child: child) : child;
  }
}
