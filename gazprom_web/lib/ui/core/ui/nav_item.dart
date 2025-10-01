import '../../../../../routing/routers.dart';

enum NavItem {
  work(Routers.work, 'assets/nav_icons/work.svg', 'Работа'),
  risk(Routers.risk, 'assets/nav_icons/risks.svg', 'Риск'),
  home(Routers.home, 'assets/nav_icons/home.svg', 'Главная'),
  news(Routers.news, 'assets/nav_icons/news.svg', 'Новости'),
  life(Routers.life, 'assets/nav_icons/life.svg', 'Жизнь');

  final String route;
  final String assetPath;
  final String label;

  const NavItem(this.route, this.assetPath, this.label);

  static final _routeMap = {for (var item in NavItem.values) item.route: item};

  static NavItem? fromRoute(String route) => _routeMap[route];
}