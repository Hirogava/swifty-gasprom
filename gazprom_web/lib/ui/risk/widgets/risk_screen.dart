import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/core/ui/button_widget.dart';
import 'package:gazprom_web/ui/core/ui/universal_container.dart';

import '../../core/themes/colors.dart';
import '../../core/themes/theme.dart';
import '../../core/ui/svg_picture_widget.dart';

import 'dart:math';

class RiskScreen extends StatelessWidget {
  const RiskScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.transparent,
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 40),
          child: Column(
            children: [
              Center(child: Text('Риски', style: AppTextStyles.style3)),
              SizedBox(height: 10),
              UniversalContainer(
                width: double.infinity,
                borderRadius: 34,
                gradient: AppColors.blackGradient3,
                children: [CardContext()],
              ),
              SizedBox(height: 10),
              _buildStockMarket(),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildStockMarket() {
    return SingleChildScrollView(
      child: Column(
        children: [
          Text('Фондовый рынок', style: AppTextStyles.style5),
          SizedBox(height: 10),
          Text('Криптовалютные инвестиции', style: AppTextStyles.style5),
          SizedBox(height: 5),
          SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Row(
              children: [
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/crypto_step.svg',
                      width: 99,
                      height: 84,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('CryptoStep', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/meme_blast.svg',
                      width: 86,
                      height: 86,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('MemeBlast', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/ico_star.svg',
                      width: 55,
                      height: 90,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('ICOStar', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/wallet_grow.svg',
                      width: 95,
                      height: 87,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('WalletGrow', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/star_meme_coin.svg',
                      width: 87,
                      height: 94,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('StarMeme Coin', style: AppTextStyles.style7),
                  ],
                ),
              ],
            ),
          ),
          SizedBox(height: 10),
          Text('Акции и фондовый рынок', style: AppTextStyles.style5),
          SizedBox(height: 5),
          SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Row(
              children: [
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/tesla_re.svg',
                      width: 89,
                      height: 89,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('TeslaRe', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/game_pop.svg',
                      width: 103,
                      height: 103,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('GamePop', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/tech_rise.svg',
                      width: 96,
                      height: 96,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('TechRise', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/bio_lip.svg',
                      width: 65,
                      height: 94,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('BioLip', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/eco_power.svg',
                      width: 58,
                      height: 92,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('EcoPower', style: AppTextStyles.style7),
                  ],
                ),
              ],
            ),
          ),
          SizedBox(height: 10),
          Text('Азартные и спекулятивные ставки', style: AppTextStyles.style5),
          SizedBox(height: 5),
          SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Row(
              children: [
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/bet_pro.svg',
                      width: 83,
                      height: 82,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('BetPro', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/lucky_draw.svg',
                      width: 80,
                      height: 77,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('LuckyDraw', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/quick_trade.svg',
                      width: 85,
                      height: 87,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('QuickTrade', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/pocker_ace.svg',
                      width: 94,
                      height: 70,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('PockerAce', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/casino_king.svg',
                      width: 117,
                      height: 62,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('CasinoKing', style: AppTextStyles.style7),
                  ],
                ),
              ],
            ),
          ),
          SizedBox(height: 10),
          Text('Сомнительные проекты', style: AppTextStyles.style5),
          SizedBox(height: 5),
          SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Row(
              children: [
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/hyip_fast_track.svg',
                      width: 96,
                      height: 86,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('HYIP-FastTrack', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/crypto_revive.svg',
                      width: 82,
                      height: 82,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('CryptoRevive', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/pyramid_fall.svg',
                      width: 63,
                      height: 73,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('daPyramidFallta', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/team_doubt.svg',
                      width: 105,
                      height: 59,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('TeamDoubt', style: AppTextStyles.style7),
                  ],
                ),
                SizedBox(width: 10),
                UniversalContainer(
                  width: 170,
                  height: 170,
                  borderRadius: 15,
                  gradient: AppColors.blackGradient3,
                  children: [
                    SvgPictureWidget(
                      path: 'assets/stock_market_icons/fakel_co.svg',
                      width: 72,
                      height: 85,
                      color: AppColors.white,
                    ),
                    Spacer(),
                    Text('FakelCo', style: AppTextStyles.style7),
                  ],
                ),
              ],
            ),
          ),
          FlipCard(
            buildFront: UniversalContainer(
              width: 170,
              height: 170,
              borderRadius: 15,
              gradient: AppColors.blackGradient3,
              children: [
                SvgPictureWidget(
                  path: 'assets/stock_market_icons/game_pop.svg',
                  width: 103,
                  height: 103,
                  color: AppColors.white,
                ),
                Spacer(),
                Text('GamePop', style: AppTextStyles.style7),
              ],
            ),
            buildBack: UniversalContainer(
              width: 170,
              height: 170,
              borderRadius: 15,
              gradient: AppColors.blackGradient3,
              children: [
                SvgPictureWidget(
                  path: 'assets/stock_market_icons/game_pop.svg',
                  width: 103,
                  height: 103,
                  color: AppColors.white,
                ),
                Spacer(),
                Text('GamePop', style: AppTextStyles.style7),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

class CardContext extends StatelessWidget {
  const CardContext({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('Мой портфель', style: AppTextStyles.style1),
        SizedBox(height: 10),
        Text('Стоимость портфеля', style: AppTextStyles.style5),
        Text('0,00', style: AppTextStyles.style3),
        Text('+ 0,00 (0.00%) за все время', style: AppTextStyles.style6),
        SizedBox(height: 10),
        Text('Диведентная дохожность в месяц', style: AppTextStyles.style5),
        Text('0,00', style: AppTextStyles.style3),
      ],
    );
  }
}

class FlipCard extends StatefulWidget {
  final Widget buildFront;
  final Widget buildBack;
  const FlipCard({
    super.key,
    required this.buildFront,
    required this.buildBack,
  });

  @override
  State<FlipCard> createState() => _FlipCardState();
}

class _FlipCardState extends State<FlipCard>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  bool _isFront = true;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      duration: const Duration(milliseconds: 600),
      vsync: this,
    );
  }

  void _flip() {
    if (_isFront) {
      _controller.forward();
    } else {
      _controller.reverse();
    }
    _isFront = !_isFront;
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: _flip,
      child: AnimatedBuilder(
        animation: _controller,
        builder: (context, child) {
          final angle = _controller.value * pi;
          final isFront = angle < pi / 2;

          return Transform(
            transform: Matrix4.identity()
              ..setEntry(3, 2, 0.001)
              ..rotateY(angle),
            alignment: Alignment.center,
            child: isFront
                ? widget.buildFront
                : Transform(
                    transform: Matrix4.identity()..rotateY(pi),
                    alignment: Alignment.center,
                    child: widget.buildBack,
                  ),
          );
        },
      ),
    );
  }

  Widget _buildFront(String path, double width, double height, String title) {
    return UniversalContainer(
      width: 170,
      height: 170,
      borderRadius: 15,
      gradient: AppColors.blackGradient3,
      children: [
        SvgPictureWidget(
          path: path,
          width: width,
          height: height,
          color: AppColors.white,
        ),
        Spacer(),
        Text(title, style: AppTextStyles.style7),
      ],
    );
  }

  Widget _buildBack(String title, String description) {
    return UniversalContainer(
      width: 170,
      height: 170,
      borderRadius: 15,
      gradient: AppColors.blackGradient3,
      children: [Text(title, style: TextStyle(color: AppColors.white)), Text(description), ButtonWidget(buttonText: 'Купить', textColor: AppColors.white, borderColor: AppColors.white, width: 59, height: 20)],
    );
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }
}
