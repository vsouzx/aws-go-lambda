resource "aws_db_instance" "default" {
  identifier           = "mydb-instance"
  allocated_storage    = 10
  db_name              = "mydb"
  engine               = "mysql"
  engine_version       = "8.0.41"
  instance_class       = "db.t3.micro"
  username             = local.db_secret.db_username
  password             = local.db_secret.db_password
  parameter_group_name = "default.mysql8.0"
  skip_final_snapshot  = true
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  publicly_accessible = true
}