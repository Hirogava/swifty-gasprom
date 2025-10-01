enum Position {
  developer('it', 'Junior-разработчик', 1, 50000),
  sisAdmin('it', 'Системный админимстратор', 1, 50000),
  analyst('it', 'Аналитик данных', 1, 50000),
  devOps('it', 'DevOps инженер', 1, 50000),
  director('it', 'Руководитель IT-проекта', 1, 50000),

  engineering('engineering', 'Инженер', 1, 350000), //temp
  medicine('medicine', 'Доктор', 1, 150000), //temp
  marketing('marketing', 'Продавец', 1, 50000), //temp
  workingProfessions('workingProfessions', 'Уборщик', 1, 250000); //temp

  final String id;
  final String title;
  final int level;
  final int salary;

  const Position(this.id, this.title, this.level, this.salary);
}
