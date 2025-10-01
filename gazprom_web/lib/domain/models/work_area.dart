enum WorkArea {
  it('it', 'IT и Технологии'),
  engineering('engineering', 'Инженерия и производство'),
  medicine('medicine', 'Медицина и здравоохранение'),
  marketing('marketing', 'Маркетинг и продажи'),
  workingProfessions('workingProfessions', 'Рабочие профессии');

  final String id;
  final String title;

  const WorkArea(this.id, this.title);
}