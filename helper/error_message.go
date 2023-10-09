package helper

import "fmt"

func Required(attribute string) string {
	return fmt.Sprintf("Maaf, %s tidak boleh kosong", attribute)
}

func GlobalError(message interface{}) string {
	if message != nil {
		return fmt.Sprintf(message.(string))
	}
	return "Maaf, Terjadi kesalahaan silahkan coba kembali"
}

func FailedToGetData(message interface{}) string {
	text := "Maaf, Gagal untuk mendapatkan data"
	if message == nil {
		return text
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func FailedToCreateData(message interface{}) string {
	text := "Maaf, Gagal untuk membuat data"
	if message == nil {
		return text
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func FailedToDeleteData(message interface{}) string {
	text := "Maaf, Gagal untuk menghapus data"
	if message == nil {
		return text
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func FailedToUpdateData(message interface{}) string {
	text := "Maaf, Gagal untuk memperbaharui data"
	if message == nil {
		return text
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func DataNotFound(message string) string {
	return fmt.Sprintf("Maaf, Data %s tidak ditemukan", message)
}

func NotValid(message string) string {
	return fmt.Sprintf("Maaf, %s tidak valid", message)
}

func Existing(message string) string {
	return fmt.Sprintf("Maaf, Data %s sudah tersedia", message)
}
