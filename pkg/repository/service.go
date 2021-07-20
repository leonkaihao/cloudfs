package repository

import (
	"fmt"
	"path"

	"github.com/leonkaihao/cloudfs/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	repoRootName        = ".repo"
	dataStoreName       = "data"
	configStoreName     = "config"
	defaultDataFileName = "data.z"
)

type Service interface {
	Table(name string) Table
	Init() error
	Load() error
}

type service struct {
	fs       utils.FS
	basePath string
	db       *gorm.DB
	tables   map[string]Table
}

func New(basePath string) Service {
	svc := &service{
		basePath: basePath,
		fs:       utils.NewBaseFs(utils.NewFs(), basePath),
		tables: map[string]Table{
			ActionsTableName: nil,
		},
	}
	return svc
}

func (svc *service) Table(name string) Table {
	return svc.tables[name]
}

func (svc *service) Init() error {
	fi, err := svc.fs.Stat(".")
	if err != nil {
		return err
	}

	if err = svc.checkEnv(); err == nil {
		return fmt.Errorf("%v: %w", fi.Name(), ErrRepoAlreadyExisted)
	}

	if err = svc.initRepo(); err != nil {
		svc.clean()
		return err
	}
	return nil
}

func (svc *service) Load() error {
	return svc.loadData()
}

// helper

func (svc *service) checkEnv() error {
	if _, err := svc.fs.Stat(repoRootName); err != nil {
		return err
	}

	configPath := path.Join(repoRootName, configStoreName)
	if _, err := svc.fs.Stat(configPath); err != nil {
		return err
	}

	dataPath := path.Join(repoRootName, dataStoreName)
	if _, err := svc.fs.Stat(dataPath); err != nil {
		return err
	}

	return nil
}

func (svc *service) initRepo() error {
	var err error
	if err = svc.fs.Mkdir(repoRootName, 0755); err != nil {
		return err
	}
	if err = svc.createConfig(); err != nil {
		return err
	}
	if err = svc.createData(); err != nil {
		return err
	}
	return nil
}

func (svc *service) createConfig() error {
	var err error

	configFs := utils.NewBaseFs(svc.fs, path.Join(repoRootName, configStoreName))
	if err = configFs.Mkdir(".", 0755); err != nil {
		return err
	}
	// Create config file

	fmt.Println("Config created.")
	return nil
}

func (svc *service) createData() error {
	var err error
	dataFs := utils.NewBaseFs(svc.fs, path.Join(repoRootName, dataStoreName))
	if err = dataFs.Mkdir(".", 0755); err != nil {
		return err
	}
	// Create data file
	if err = svc.loadData(); err != nil {
		return err
	}

	fmt.Println("Data created.")
	return nil
}

func (svc *service) clean() {
	svc.fs.RemoveAll(repoRootName)
	fmt.Println("Repo cleaned.")
}

func (svc *service) loadData() error {
	if svc.db != nil {
		return nil
	}
	dataFilePath := path.Join(svc.basePath, repoRootName, dataStoreName, defaultDataFileName)
	db, err := gorm.Open(sqlite.Open(dataFilePath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("%v: %w", dataFilePath, err)
	}
	svc.db = db
	svc.tables[ActionsTableName] = NewActions(db)
	return nil
}
