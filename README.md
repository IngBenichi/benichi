# Benichi 🚀

Benichi es un **lenguaje de programación interpretado** creado en Go.  
Tiene su propio REPL, permite ejecutar archivos `.bn` y está diseñado para ser fácil de usar y extender.

---

## 🔹 Características

- REPL interactivo (`benichi`) al estilo Python.
- Ejecutar archivos `.bn` (`benichi archivo.bn`).
- Instrucción `print` para mostrar texto en consola.
- Listo para expandir con variables, funciones y módulos.
- Compilado en Go, sin depender de Python.

---

## 🔹 Instalación

### Windows
1. Descarga el binario desde [Releases](https://github.com/IngBenichi/benichi/releases).
2. Coloca `benichi.exe` en una carpeta de tu elección, por ejemplo `C:\benichi`.
3. Agrega la carpeta al **PATH** de Windows.
4. Abre CMD o PowerShell y prueba:

```bash
benichi --help
````

### Linux / Mac

1. Descarga el binario correspondiente.
2. Dale permisos de ejecución:

```bash
chmod +x benichi
```

3. Mueve el binario a una carpeta en tu PATH, por ejemplo `/usr/local/bin`.
4. Prueba:

```bash
benichi --help
```

---

## 🔹 Uso

### Ejecutar archivo `.bn`

```bash
benichi main.bn
```

### Abrir REPL interactivo

```bash
benichi
```

Luego podrás escribir comandos como:

```bn
>>> print "Hola mundo"
Hola mundo
>>> exit
```

---

## 🔹 Ejemplo de archivo `.bn`

```bn
# main.bn
print "Hola mundo desde Benichi 🚀"
```

---

## 🔹 Contribuir

Si quieres contribuir a Benichi:

1. Haz un fork del repo.
2. Crea tu branch: `git checkout -b mi-nueva-funcionalidad`.
3. Haz commit de tus cambios: `git commit -m "Agrega nueva funcionalidad"`.
4. Haz push: `git push origin mi-nueva-funcionalidad`.
5. Abre un Pull Request.

---
