docker exec -it harmony-db bash -c "\
until psql -h localhost -U postgres -d harmony-db -c 'select 1'>/dev/null 2>&1;\
do\
  echo 'Waiting for postgres server....';\
  sleep 1;\
done;\
exit;\
"
echo "DB Connected !!"
echo "Running Migrations..."
liquibase --username=postgres --password=password --url=jdbc:postgresql://localhost:5432/harmony-db --changeLogFile=database/changelog.xml update

