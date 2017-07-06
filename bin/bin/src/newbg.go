package main
import (
	"./yml"
    "./api/unsplash"
	"flag"
    "github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	"image/draw"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type RGBA struct {
	R, G, B, A uint8
}

func main() {

	//FLAGS
	e := flag.String("e", "", "configuration and env yaml location")
	flag.Parse()

	//GET YML ENV
	y := yml.New(*e)

	//GET JSON
	k := y.Env["unsplash_key"]
	u := unsplash.New(k)

	//GET SAVE DIR
	f := y.Env["bg_location"]
	b := y.Env["bg_backup"]
    c := y.Env["bg_cmd"]
    a := y.Env["bg_arg"]
    a1 := y.Env["bg_arg1"]

	//DOWNLOAD FILE
	DownloadFile(f, u.Urls["raw"])

	//SET ENVS
	y.Set("bg_url", u.Links["html"])
	y.Set("dyn_color", u.Color)
	y.Set("bg_name", u.Id)
	y.Write(*e)

	SetBG(c, a, a1, f, b, u)
}

func GenColor(h string)(c color.RGBA) {
    cl, err := colorful.Hex(h)
    r, g, b := cl.FastLinearRgb()
    c = color.RGBA{uint8(r), uint8(g), uint8(b), 1}
    if err != nil{
        log.Fatal(err)
    }
    return
}

func SetBG(c, a, a1, f, b string, u *unsplash.Unsplash) {

	bg := exec.Command(c, a, a1, f)
	// Backup wallpaper ----------------------
	b = b + "/" + u.Id
	cp := exec.Command("cp", f, b)

	// Run Commands -----------------------
	err := cp.Run()
	err = bg.Run()
	if err != nil {
	}
}

func Border(i image.Image, fg, bg color.Color, d int) (o image.Image) {
	x, y := GetImageDimensions(i)
    s := 12
	b := i.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, i, b.Min, draw.Src)
	for cx := 0; cx < x; cx = cx + 1 {
		for cy := 0; cy < y; cy = cy + 1 {
			if cy <= (d + s) || cy >= (y-d-s) {
                m.Set(cx, cy, fg)
            }
			if cx <= (d + s) || cx >= (x-d-s) {
                m.Set(cx, cy, fg)
            }

            if cy <= d || cy >= (y-d) {
				m.Set(cx, cy, bg)
			}
			if cx <= d || cx >= (x-d) {
				m.Set(cx, cy, bg)
			}
		}
	}
	o = m
	return
}

func GetImageDimensions(img image.Image) (image_x, image_y int) {
	b := img.Bounds()
	image_x = b.Max.X
	image_y = b.Max.Y
	return
}

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
