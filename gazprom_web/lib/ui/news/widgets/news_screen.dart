import 'package:flutter/material.dart';
import 'package:gazprom_web/ui/core/ui/svg_picture_widget.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../core/themes/colors.dart';

class NewsScreen extends StatelessWidget {
  const NewsScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Background(
      child: Scaffold(
        appBar: AppBar(
          title: Text(
            'Новости',
            style: GoogleFonts.roboto(fontSize: 24, color: AppColors.white),
          ),
          backgroundColor: AppColors.transparent,
          shape: const Border(bottom: BorderSide(color: Colors.white)),
        ),
        backgroundColor: AppColors.transparent,
        body: SingleChildScrollView(
              child: Padding(
                padding: EdgeInsets.symmetric(horizontal: 16, vertical: 20),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            'Массовая закупка баранов произошла в Чечне. Рост спроса стимулирует рост цен в секторе “Сомнительные Проекты”',
                            style: GoogleFonts.roboto(
                              fontSize: 16,
                              color: AppColors.white,
                            ),
                            softWrap: true,
                            overflow: TextOverflow.visible,
                          ),
                        ),
                        SizedBox(width: 8),
                        SvgPictureWidget(
                          path: 'assets/news_icons/1.svg',
                          width: 82,
                          height: 82,
                          color: AppColors.white,
                        ),
                      ],
                    ),
                    Divider(color: AppColors.white),
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            'Массовая закупка баранов произошла в Чечне. Рост спроса стимулирует рост цен в секторе “Сомнительные Проекты”',
                            style: GoogleFonts.roboto(
                              fontSize: 16,
                              color: AppColors.white,
                            ),
                          ),
                        ),
                        SizedBox(width: 8),
                        SvgPictureWidget(
                          path: 'assets/news_icons/2.svg',
                          width: 68,
                          height: 86,
                          color: AppColors.white,
                        ),
                      ],
                    ),
                    Divider(color: AppColors.white),
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            'Массовая закупка баранов произошла в Чечне. Рост спроса стимулирует рост цен в секторе “Сомнительные Проекты”',
                            style: GoogleFonts.roboto(
                              fontSize: 16,
                              color: AppColors.white,
                            ),
                          ),
                        ),
                        SizedBox(width: 8),
                        SvgPictureWidget(
                          path: 'assets/news_icons/3.svg',
                          width: 86,
                          height: 80,
                          color: AppColors.white,
                        ),
                      ],
                    ),
                    Divider(color: AppColors.white),
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            'Массовая закупка баранов произошла в Чечне. Рост спроса стимулирует рост цен в секторе “Сомнительные Проекты”',
                            style: GoogleFonts.roboto(
                              fontSize: 16,
                              color: AppColors.white,
                            ),
                          ),
                        ),
                        SizedBox(width: 8),
                        SvgPictureWidget(
                          path: 'assets/news_icons/4.svg',
                          width: 86,
                          height: 50,
                          color: AppColors.white,
                        ),
                      ],
                    ),
                    Divider(color: AppColors.white),
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            'Массовая закупка баранов произошла в Чечне. Рост спроса стимулирует рост цен в секторе “Сомнительные Проекты”',
                            style: GoogleFonts.roboto(
                              fontSize: 16,
                              color: AppColors.white,
                            ),
                          ),
                        ),
                        SizedBox(width: 8),
                        SvgPictureWidget(
                          path: 'assets/news_icons/5.svg',
                          width: 84,
                          height: 83,
                          color: AppColors.white,
                        ),
                      ],
                    ),
                    Divider(color: AppColors.white),
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            'Массовая закупка баранов произошла в Чечне. Рост спроса стимулирует рост цен в секторе “Сомнительные Проекты”',
                            style: GoogleFonts.roboto(
                              fontSize: 16,
                              color: AppColors.white,
                            ),
                          ),
                        ),
                        SizedBox(width: 8),
                        SvgPictureWidget(
                          path: 'assets/news_icons/6.svg',
                          width: 61,
                          height: 80,
                          color: AppColors.white,
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
      ),
    );
  }
}
