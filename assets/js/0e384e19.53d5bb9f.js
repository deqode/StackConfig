(self.webpackChunkappconfig=self.webpackChunkappconfig||[]).push([[671],{426:function(e,n,i){"use strict";i.r(n),i.d(n,{frontMatter:function(){return p},contentTitle:function(){return l},metadata:function(){return s},toc:function(){return c},default:function(){return g}});var t=i(2122),a=i(9756),o=(i(7294),i(3905)),r=["components"],p={},l="What is App Config",s={unversionedId:"intro",id:"intro",isDocsHomePage:!1,title:"What is App Config",description:"About",source:"@site/docs/intro.md",sourceDirName:".",slug:"/intro",permalink:"/StackConfig/docs/intro",editUrl:"https://github.com/facebook/docusaurus/edit/master/website/docs/intro.md",version:"current",frontMatter:{},sidebar:"tutorialSidebar"},c=[{value:"About",id:"about",children:[]},{value:"Where can be used?",id:"where-can-be-used",children:[]},{value:"Usage",id:"usage",children:[{value:"Example with leveldb and zap logger",id:"example-with-leveldb-and-zap-logger",children:[]},{value:"App service Interface",id:"app-service-interface",children:[]}]}],u={toc:c};function g(e){var n=e.components,i=(0,a.Z)(e,r);return(0,o.kt)("wrapper",(0,t.Z)({},u,i,{components:n,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"what-is-app-config"},"What is App Config"),(0,o.kt)("h2",{id:"about"},"About"),(0,o.kt)("p",null,"This is a Golang library that provides service to store key value for an application configuration. It uses Gokv key value store for storing multiple version of an application config. An application might consist of multiple services, an array of such services is stored in the application configuration. Each service contains details such as service runtime, databse, CPU, memory, network, scaling and other general settings including environment variables."),(0,o.kt)("h2",{id:"where-can-be-used"},"Where can be used?"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"If you want to maintain multiple versions of an application configuration"),(0,o.kt)("li",{parentName:"ul"},"If you want to maintain multiple services configuration in an application configuration"),(0,o.kt)("li",{parentName:"ul"},"If you want to generate cloud infrastructure template for an app")),(0,o.kt)("h2",{id:"usage"},"Usage"),(0,o.kt)("p",null,"go get github.com/deqodelabs/IaaC"),(0,o.kt)("h3",{id:"example-with-leveldb-and-zap-logger"},"Example with leveldb and zap logger"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'import (\n    "github.com/deqodelabs/IaaC/appconfig"\n    "github.com/philippgille/gokv/leveldb"\n    "go.uber.org/zap"\n)\n\noptions := leveldb.DefaultOptions\nstore, err := leveldb.NewStore(options)\nif err != nil {\n    panic(err)\n}\nlogger := zap.NewExample()\nappService := appconfig.AppService{\n    Store:  store,\n    Logger: logger,\n}\n')),(0,o.kt)("h3",{id:"app-service-interface"},"App service Interface"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// used to validate an app config\nValidateAppConfig(app *pb.AppConfig) error\n// save latest version of app config to key-value store\nSave(app *pb.AppConfig) (*pb.AppConfig, error)\n// get latest version of app config from key-value store\nGetAppConfig(id string) (*pb.AppConfig, error)\n// get app config corresponding to any available version from key-value store\nGetAppConfigForVersion(id string, version int32) (*pb.AppConfig, error)\n")))}g.isMDXComponent=!0}}]);