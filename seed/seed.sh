cd ../database/schema
numOfMigrations=$(ls | wc -l)
for i in $( eval echo {1..$numOfMigrations} )
do
goose postgres postgres://root:root@127.0.0.1:5432/restaurant_management_system?sslmode=disable down
done
goose postgres postgres://root:root@127.0.0.1:5432/restaurant_management_system?sslmode=disable up
cd ../../seed
go run seedAdmin.go