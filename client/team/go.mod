module muxi-workbench-team-client

replace muxi-workbench => ../../

replace muxi-workbench-feed => ../../service/feed

replace muxi-workbench-project => ../../service/project

replace muxi-workbench-user => ../../service/user

replace muxi-workbench-team => ../../service/team

go 1.14

require github.com/micro/go-micro v1.18.0 // indirect
