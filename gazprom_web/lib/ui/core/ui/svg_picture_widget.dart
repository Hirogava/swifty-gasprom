import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

class SvgPictureWidget extends StatelessWidget {
  final String path;
  final double width;
  final double height;
  final Color color;
  const SvgPictureWidget({super.key, required this.path, required this.width, required this.height, required this.color});

  @override
  Widget build(BuildContext context) {
    return SvgPicture.asset(
      path,
      width: width,
      height: height,
      colorFilter: ColorFilter.mode(color, BlendMode.srcIn),
    );
  }
}
