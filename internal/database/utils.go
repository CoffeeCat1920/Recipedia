package database

func (s *service) columnExist(tableName string, columnName string) (bool, error) {

  query := `
  SELECT EXISTS (
    SELECT 1 
    FROM information_schema.columns 
    WHERE table_schema = 'public'
    AND table_name = $1 
    AND column_name = $2
  );`

  var exists bool
  err := s.db.QueryRow(query, tableName, columnName).Scan(&exists)
  if err != nil {
    return false, err
  }

  return exists, nil
  
}
