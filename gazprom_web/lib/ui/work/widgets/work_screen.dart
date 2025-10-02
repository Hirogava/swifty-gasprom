import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:gazprom_web/providers/current_user_provider.dart';
import 'package:gazprom_web/ui/core/themes/theme.dart';
import 'package:gazprom_web/ui/core/ui/universal_container.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../domain/models/position.dart';
import '../../../domain/models/user.dart';
import '../../../domain/models/work_area.dart';
import '../../core/themes/colors.dart';

class WorkScreen extends ConsumerStatefulWidget {
  const WorkScreen({super.key});

  @override
  ConsumerState<WorkScreen> createState() => _WorkScreenState();
}

class _WorkScreenState extends ConsumerState<WorkScreen> {
  WorkArea? _selectedArea;
  bool _choosing = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
          title: Text(
            'Работа',
            style: GoogleFonts.roboto(fontSize: 24, color: AppColors.white),
          ),
          backgroundColor: AppColors.transparent,
          shape: const Border(bottom: BorderSide(color: Colors.white)),
        ),
      backgroundColor: AppColors.transparent,
      body: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 40),
        child: !_choosing
            ? _buildStatus(ref)
            : (_selectedArea == null
                  ? _buildWorkAreasView()
                  : _buildPositionsView(_selectedArea!)),
      ),
    );
  }

  Widget _buildStatus(WidgetRef ref) {
    final user = ref.watch(currentUserProvider);

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Профессия: ${user?.position?.title ?? 'Безработный'}',
          style: AppTextStyles.style3,
        ),
        const SizedBox(height: 10),
        Text(
          'Уровень: ${user?.position?.level ?? 0}',
          style: AppTextStyles.style3,
        ),
        const SizedBox(height: 10),
        Text(
          'Зарплата: ${user?.position?.salary ?? 0} руб',
          style: AppTextStyles.style3,
        ),
        const SizedBox(height: 10),
        Text('Счастье: ${user?.happiness ?? 0}', style: AppTextStyles.style3),
        const SizedBox(height: 20),

        UniversalContainer(
          borderRadius: 15,
          isButton: true,
          onTap: () {
            setState(() {
              _choosing = true;
              _selectedArea = null;
            });
          },
          children: [
            Text(
              user?.position == null ? 'Выбрать работу' : 'Сменить работу',
              style: AppTextStyles.style3,
            ),
          ],
        ),
      ],
    );
  }

  Widget _buildWorkAreasView() {
    return Column(
      children: [
        Align(
          alignment: Alignment.centerLeft,
          child: TextButton.icon(
            onPressed: () {
              setState(() {
                _choosing = false;
              });
            },
            icon: const Icon(Icons.arrow_back, color: Colors.white),
            label: const Text("Назад", style: AppTextStyles.style3),
          ),
        ),
        const SizedBox(height: 16),
        Expanded(
          child: ListView.separated(
            itemBuilder: (context, i) {
              final area = WorkArea.values[i];
              return UniversalContainer(
                borderRadius: 15,
                isButton: true,
                onTap: () {
                  setState(() {
                    _selectedArea = area;
                  });
                },
                children: [Text(area.title, style: AppTextStyles.style3)],
              );
            },
            separatorBuilder: (_, __) => const SizedBox(height: 10),
            itemCount: WorkArea.values.length,
          ),
        ),
      ],
    );
  }

  Widget _buildPositionsView(WorkArea area) {
    final positions = Position.values.where((p) => p.id == area.id).toList();

    return Column(
      children: [
        Align(
          alignment: Alignment.centerLeft,
          child: TextButton.icon(
            onPressed: () {
              setState(() {
                _selectedArea = null;
              });
            },
            icon: const Icon(Icons.arrow_back, color: Colors.white),
            label: const Text("Назад", style: AppTextStyles.style3),
          ),
        ),
        const SizedBox(height: 16),
        Expanded(
          child: ListView.separated(
            itemBuilder: (context, i) {
              final pos = positions[i];
              return UniversalContainer(
                borderRadius: 15,
                isButton: true,
                onTap: () {
                  ref.read(currentUserProvider.notifier).update((user) {
                    if (user == null) return null;
                    return User(
                      id: user.id,
                      bankUserId: user.bankUserId,
                      accessToken: user.accessToken,
                      refreshToken: user.refreshToken,
                      position: pos,
                      happiness: user.happiness,
                    );
                  });

                  setState(() {
                    _choosing = false;
                    _selectedArea = null;
                  });
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Вы выбрали: ${pos.title}')),
                  );
                },
                children: [
                  Text(pos.title, style: AppTextStyles.style3),
                  Text(
                    "Зарплата: ${pos.salary} ₽",
                    style: AppTextStyles.style3,
                  ),
                ],
              );
            },
            separatorBuilder: (_, __) => const SizedBox(height: 10),
            itemCount: positions.length,
          ),
        ),
      ],
    );
  }
}
