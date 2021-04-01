package appconfig

import (
	"errors"
	"github.com/deqodelabs/IaaC/appconfig/pb"
	"github.com/deqodelabs/IaaC/appconfig/tools"
	"github.com/deqodelabs/IaaC/appconfig/validate"
	"github.com/nu7hatch/gouuid"
	"github.com/philippgille/gokv"
	"go.uber.org/zap"
)

type AppService	struct{
	Store gokv.Store
	Logger zap.Logger
}

type VersionRange struct{
	LastVersion int32
	LatestVersion int32
}

func(s *AppService) ValidateAppConfig(app *pb.AppConfig) error {
	s.Logger.Debug("validate app config method called with:", zap.Any("app", app))
	err := app.Validate()
	if err != nil{
		s.Logger.Error("error:", zap.Error(err))
		return err
	}
	err = validate.CustomValidate(app)
	if err != nil{
		s.Logger.Error("error:", zap.Error(err))
		return err
	}
	return nil
}

func (s *AppService) Save(app *pb.AppConfig) (*pb.AppConfig, error){
	s.Logger.Debug("save app config method called with:", zap.Any("app", app))
	err := s.ValidateAppConfig(app)
	if err != nil{
		s.Logger.Error("error:", zap.Error(err))
		return app, err
	}
	var versionRange VersionRange
	var versionId string
	if app.AppId == "" {
		id, _ := uuid.NewV4()
		app.AppId = id.String()
		versionRange = VersionRange{1,1}
		versionId = tools.GenerateVersionId(app.AppId, 1)
	}else{
		found, err := s.Store.Get(app.AppId, versionRange)
		if err != nil {
			s.Logger.Error("error:", zap.Error(err))
			return app, err
		}
		if !found {
			s.Logger.Error("error: version range not found")
			return app, errors.New("version range not found")
		}
		latestVersion := versionRange.LatestVersion
		latestVersion++
		if app.Version < latestVersion {
			s.Logger.Error("error: version already exist")
			return app, errors.New("version already exist")
		}else if app.Version > latestVersion {
			s.Logger.Error("error: version missed")
			return app, errors.New("version missed")
		}else{
			versionRange.LatestVersion++
		}
		versionId = tools.GenerateVersionId(app.AppId, latestVersion)
	}
	err = s.Store.Set(app.AppId, versionRange)
	if err != nil {
		s.Logger.Error("error:", zap.Error(err))
		return app, err
	}
	err = s.Store.Set(versionId, app)
	if err != nil {
		s.Logger.Error("error:", zap.Error(err))
		return app, err
	}
	return app, nil
}

func (s *AppService) GetAppConfig(id string) (*pb.AppConfig, error){
	s.Logger.Debug("get app config method called with:", zap.String("id", id))
	var versionRange VersionRange
	var appConfig *pb.AppConfig
	found, err := s.Store.Get(id, versionRange)
	if err != nil {
		s.Logger.Error("error:", zap.Error(err))
		return nil, err
	}
	if !found {
		s.Logger.Error("error: id for app config not found")
		return nil, errors.New("id for app config not found")
	}
	versionId := tools.GenerateVersionId(id, versionRange.LatestVersion)
	found, err = s.Store.Get(versionId, appConfig)
	if err != nil {
		s.Logger.Error("error:", zap.Error(err))
		return nil, err
	}
	if !found {
		s.Logger.Error("error: app config for version not found")
		panic("app config for version not found")
	}
	return appConfig, nil
}

func (s *AppService) GetAppConfigForVersion(id string, version int32) (*pb.AppConfig, error){
	s.Logger.Debug("get app config for version method called with:", zap.String("id", id), zap.Int32("version", version))
	var versionRange VersionRange
	var appConfig *pb.AppConfig
	found, err := s.Store.Get(id, versionRange)
	if err != nil {
		s.Logger.Error("error:", zap.Error(err))
		return nil, err
	}
	if !found {
		s.Logger.Error("error: id for app config not found")
		return nil, errors.New("id for app config not found")
	}
	if version < versionRange.LastVersion {
		s.Logger.Error("error: version is deprecated")
		return nil, errors.New("version is deprecated")
	}else if version > versionRange.LatestVersion {
		s.Logger.Error("error: version does not exist")
		return nil, errors.New("version does not exist")
	}
	versionId := tools.GenerateVersionId(id, version)
	found, err = s.Store.Get(versionId, appConfig)
	if err != nil {
		s.Logger.Error("error:", zap.Error(err))
		return nil, err
	}
	if !found {
		s.Logger.Error("error: app config for version not found")
		panic("app config for version not found")
	}
	return appConfig, nil
}

