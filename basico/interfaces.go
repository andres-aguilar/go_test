package main

import "fmt"

type User interface {
  Permisos() int
  Nombre() string
}

// implementando la interfaz User en la estructura Admin
type Admin struct {
  nombre string
}
func (this Admin) Permisos() int {
  return 5
}
func (this Admin) Nombre() string {
  return this.nombre
}


/* Funcion para autenticar
 * recibe un dato de tipo User
 */
func auth(user User) string {
  permisos := user.Permisos()

  if permisos == 5 {
    return user.Nombre() + " es un administrador"
  }else if permisos == 4 {
    return user.Nombre() + " es moderador"
  }
  return user.Nombre() + " no tiene permisos especiales"
}

func main()  {
  admin := Admin{"Andres"}

  fmt.Println(auth(admin))
}
