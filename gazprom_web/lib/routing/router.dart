import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/life/widgets/life_screen.dart';
import 'package:gazprom_web/ui/news/widgets/news_screen.dart';
import 'package:gazprom_web/ui/risk/widgets/risk_screen.dart';
import 'package:gazprom_web/ui/work/widgets/work_screen.dart';
import 'package:go_router/go_router.dart';

import '../ui/core/ui/custom_navigation_bar.dart';
import '../ui/home/widgets/home_screen.dart';
import 'routers.dart';


final GoRouter router = GoRouter(
  routes: <RouteBase>[
    ShellRoute(
      builder: (BuildContext context, GoRouterState state, Widget child) {
        return CustomNavigationBar(child: child);
      },
      routes: [
        GoRoute(
          path: Routers.work,
          builder: (BuildContext context, GoRouterState state) {
            return const WorkScreen();
          },
        ),
        GoRoute(
          path: Routers.risk,
          builder: (BuildContext context, GoRouterState state) {
            return const RiskScreen();
          },
        ),
        GoRoute(
          path: Routers.home,
          builder: (BuildContext context, GoRouterState state) {
            return const HomeScreen();
          },
        ),
        GoRoute(
          path: Routers.news,
          builder: (BuildContext context, GoRouterState state) {
            return const NewsScreen();
          },
        ),
        GoRoute(
          path: Routers.life,
          builder: (BuildContext context, GoRouterState state) {
            return const LifeScreen();
          },
        ),
      ],
    ),
  ],
);
