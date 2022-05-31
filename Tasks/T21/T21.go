package main

import "fmt"

type AnalytycalDataService interface { //интерфейс внешнего сервиса
	SendXmlData()
}

type XmlDocument struct { //структура внешнего сервиса
}

func (d XmlDocument) SendXmlData() { //метод внешнего сервиса
	fmt.Println("Sendind XML")
}

type JsonDocument struct { //внутренняя структура
}

func (d JsonDocument) ConvertToXml() string { //метод внутренней структуры
	return "<xml></xml>"
}

type JsonDocumentAdapter struct { //адаптер
	JsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() { //метод адаптера
	adapter.JsonDocument.ConvertToXml()
	fmt.Println("Sendind XML converted")
}

/*вместо того, чтобы создавать новый метод для внутренней структуры
мы встраиваем ее в адаптер, для которого реализуем необходимый метод. Это
позволяет скрыть преобразования от клиента, который получает доступ к адаптеру
*/

func main() {
	var jsonDoc JsonDocument

	adapter := JsonDocumentAdapter{
		JsonDocument: &jsonDoc,
	}
	adapter.SendXmlData()
}
