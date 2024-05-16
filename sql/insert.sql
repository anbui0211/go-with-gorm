INSERT INTO locations 
VALUE 
  ('LC01', 'Nguyễn Công Trứ, Sơn trà, Đà Nẵng','Đà Nẵng', 'Sơn Trà', 'Nguyễn Công Trứ'),
  ('LC02', 'Nguyễn Hữu Thọ, Hải Châu, Đà Nẵng', 'Đà Nẵng', 'Hải Châu', 'Nguyễn Hữu Thọ'),
  ('LC03', 'Trưng Nữ Vương, Hải Châu, Đà Nẵng', 'Đà Nẵng', 'Hải Châu', 'Trưng Nữ Vương');

INSERT INTO users 
VALUE
  ('US01','Ân', 10, 1, 1, 'active', 'LC01'),
  ('US02', 'Định', 20, 1, 2, 'inactive', 'LC02'),
  ('US03', 'Thu', 30, 0, 2, 'inactive', 'LC01'),
  ('US04', 'Phương', 40, 1, 2, 'inactive', 'LC03'),
  ('US05', 'Huy', 50, 1, 2, 'inactive', 'LC02');