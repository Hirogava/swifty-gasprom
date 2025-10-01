import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/start/widgets/start_screen_animated.dart';
import 'package:go_router/go_router.dart';

class ButtonWidget extends StatefulWidget {
  final String buttonText;
  final Color textColor;
  final Color? backgroundColor;
  final Color borderColor;
  final LinearGradient? gradient;
  final double width;
  final double height;
  final String? pathScreen;
  final ValueChanged<Widget>? onTap;
  final VoidCallback? hm;
  const ButtonWidget({
    super.key,
    required this.buttonText,
    required this.textColor,
    this.backgroundColor,
    required this.borderColor,
    this.gradient,
    required this.width,
    required this.height,
    this.pathScreen,
    this.onTap,
    this.hm,
  });

  @override
  State<ButtonWidget> createState() => _ButtonWidgetState();
}

class _ButtonWidgetState extends State<ButtonWidget> {
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        if (widget.onTap != null) {
          widget.onTap!(const StartScreenAnimated());
        } else if (widget.pathScreen != null) {
          context.go(widget.pathScreen!);
          if (widget.hm != null) {
            widget.hm!();
          }
        }
      },
      child: Container(
        width: widget.width,
        height: widget.height,
        decoration: BoxDecoration(
          color: widget.gradient == null ? widget.backgroundColor : null,
          gradient: widget.gradient,
          borderRadius: BorderRadius.circular(15),
          border: Border.all(color: widget.borderColor),
        ),
        padding: const EdgeInsets.symmetric(vertical: 10),
        child: Center(
          child: Text(
            widget.buttonText,
            style: TextStyle(color: widget.textColor),
          ),
        ),
      ),
    );
  }
}
