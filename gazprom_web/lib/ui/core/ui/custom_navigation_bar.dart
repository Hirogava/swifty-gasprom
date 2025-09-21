import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:go_router/go_router.dart';

import '../themes/colors.dart';
import 'nav_item.dart';

class CustomNavigationBar extends StatefulWidget {
  final Widget child;

  const CustomNavigationBar({super.key, required this.child});

  @override
  State<CustomNavigationBar> createState() => _CustomNavigationBarState();
}

class _CustomNavigationBarState extends State<CustomNavigationBar> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          Positioned.fill(child: widget.child),
          Positioned(
            left: 0,
            right: 0,
            bottom: 0,
            child: Container(
              padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
              decoration: BoxDecoration(
                color: AppColors.blue2,
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(30),
                  topRight: Radius.circular(30),
                ),
                border: Border(
                  top: BorderSide(color: AppColors.white, width: 2),
                ),
              ),
              child: Row(
                children: [
                  ...NavItem.values.map((item) => Expanded(child: NavItemWidget(item: item))),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class NavItemWidget extends StatelessWidget {
  final NavItem item;

  const NavItemWidget({super.key, required this.item});

  @override
  Widget build(BuildContext context) {
    final isSelected = _isCurrentRoute(context, item.route);
    return GestureDetector(
      behavior: HitTestBehavior.opaque,
      onTap: () => _handleTap(context),
      child: Container(
        padding: const EdgeInsets.all(8),
        margin: const EdgeInsets.only(bottom: 4),
        child: Column(
          children: [
            SvgPicture.asset(
              item.assetPath,
              width: isSelected ? 40 : 24,
              height: isSelected ? 40 : 24,
              alignment: Alignment.center,
              colorFilter: ColorFilter.mode(
                isSelected ? AppColors.white : AppColors.unselectedItem,
                BlendMode.srcIn,
              ),
            ),
            
              const SizedBox(height: 8),
              Text(
                item.label,
                style: TextStyle(
                  color: isSelected
                      ? AppColors.white
                      : AppColors.unselectedItem,
                ),
              ),
            ],
          
        ),
      ),
    );
  }

  bool _isCurrentRoute(BuildContext context, String route) {
    return GoRouterState.of(context).uri.toString() == route;
  }

  void _handleTap(BuildContext context) {
    GoRouter.of(context).go(item.route);
  }
}
