package models

type CodeArea string
const(
	YaTechManager CodeArea ="tech-manager"
	YaDB CodeArea ="database-developer"
	YaDesktop CodeArea ="desktop-developer"
	YaFrontend CodeArea ="frontend-developer"
	YaBackend CodeArea="backend-developer"
	YaFullStack CodeArea="full-stack-developer"
	YaML CodeArea="ml-developer"
	YaMobile CodeArea="mob-app-developer"
	YaMobileAndroid CodeArea="mob-app-developer-android"
	YaMobileIOS CodeArea="mob-app-developer-ios"
	YaNOC CodeArea="noc-developer"
	YaSystem CodeArea="system-developer"
	YaDevOps CodeArea="dev-ops"
)

type ParamForYandex func(yandex) yandex

func SetParamAreaForYandex(area CodeArea) ParamForYandex {
	return func(ya yandex) yandex {
		ya.link.AddQueryParam("public_professions",string(area))
		return ya
	}
}

func SetParamCityForYandex(city string) ParamForYandex {
	return func(ya yandex) yandex {
		ya.link.AddQueryParam("cities",city)
		return ya
	}
}

func SetParamLevelForYandex(level string) ParamForYandex {
	return func(ya yandex) yandex {
		ya.link.AddQueryParam("pro_levels",level)
		return ya
	}
}
type yandex struct {
	link *Link
	name string
}
func InitYandex(params...ParamForYandex) *yandex {

	ya := yandex{ link: &Link{Host:  "yandex.ru",
	 								 Scheme: "https", 
	 								 Path: "jobs/api/publications"},
					name: "Yandex"}
	for _, param := range params {
		ya = param(ya)
	}	
	return &ya
}

func (ya *yandex) GetName() string {
	return ya.name
}
func (ya *yandex) GetJobs() ([]Job, error) {
	yar, err := DoRequest[YandexResponse](ya.link.GetFullURL())
	if err != nil {
		return nil, err
	}
	var jobs = yar.GetJobs()

	return jobs, nil
}
