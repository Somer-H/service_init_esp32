package adapters

import (
	"fmt"
	"log"
	"status_esp32_service/core"
	"status_esp32_service/src/esp32/domain/entities"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() (*MySQL, error) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}, nil
}

func (m *MySQL) ChangeStatus(name string, status string) (*entities.Esp32, error) {
    query1 := "SELECT * FROM esp32 where name = ?"; 
    result1, err := m.conn.FetchRows(query1, name)
    if err != nil {
        return nil, err
    }
    var esp32Get entities.Esp32
    if result1.Next() {
        err := result1.Scan(&esp32Get.ID, &esp32Get.Name, &esp32Get.Status)
        if err != nil {
            return nil, err
        }
    }else{
        return nil, fmt.Errorf("no se ha encontrado el dispositivo")
    }
    if esp32Get.Status == "activate"{
        return &esp32Get, fmt.Errorf("el dispositivo ya está activado")
    }

	var esp entities.Esp32
    query := "UPDATE esp32 SET status = ? WHERE name = ?"
    result, err := m.conn.ExecutePreparedQuery(query, status, name)
    if err != nil {
        return nil, err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return nil, err
    }

    if rowsAffected == 0 {
        return nil, fmt.Errorf("no se ha hecho ningún cambio, ha habido un error")
    }

    return &esp, nil
}