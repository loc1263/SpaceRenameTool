package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func listarArchivosConEspacios(directorioRaiz string) error {
	fmt.Println("Archivos con espacios en blanco en el nombre:")
	return filepath.Walk(directorioRaiz, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error al acceder a %s: %v\n", ruta, err)
			return nil
		}

		if !info.IsDir() && strings.Contains(info.Name(), " ") {
			fmt.Printf("%s\n", ruta)
		}
		return nil
	})
}

func renombrarArchivosConEspacios(directorioRaiz string) error {
	fmt.Println("Renombrando archivos con espacios en blanco:")
	return filepath.Walk(directorioRaiz, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error al acceder a %s: %v\n", ruta, err)
			return nil
		}

		if !info.IsDir() && strings.Contains(info.Name(), " ") {
			nuevoNombre := strings.ReplaceAll(info.Name(), " ", "_")
			nuevaRuta := filepath.Join(filepath.Dir(ruta), nuevoNombre)

			if err := os.Rename(ruta, nuevaRuta); err != nil {
				fmt.Printf("Error al renombrar %s: %v\n", ruta, err)
				return nil
			}

			fmt.Printf("Archivo renombrado: %s -> %s\n", ruta, nuevaRuta)
		}
		return nil
	})
}

func main() {

	var directorioRaiz string
	if len(os.Args) > 1 {
		directorioRaiz = os.Args[1]
	} else {
		fmt.Println("Uso: programa <directorioRaiz>")
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error al obtener el nombre de la máquina:", err)
		return
	}

	fmt.Println("El programa se está ejecutando en la máquina:", hostname)

	for {
		fmt.Println()
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Listar archivos con espacios en blanco en el nombre")
		fmt.Println("2. Renombrar archivos reemplazando espacios por _")
		fmt.Println("3. Salir")

		var opcion int
		fmt.Print("Opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Directorio raiz:", directorioRaiz)
			err := listarArchivosConEspacios(directorioRaiz)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 2:
			err := renombrarArchivosConEspacios(directorioRaiz)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			fmt.Println("Saliendo del programa.")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}
