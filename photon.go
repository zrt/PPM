package PPM

type PhotonMapping struct {
	world  World
	camera Camera
	path   string
	img    Image
}

func NewPhotonMapping(world World, camera Camera, path string) *PhotonMapping {
	pm := &PhotonMapping{world, camera, path, Image{}}
	pm.img.Init(camera.X, camera.Y)
	return pm
}

func (pm *PhotonMapping) Render() {

	// debug
	pm.img.DebugSave("debug_"+pm.path, "algo:pm, pm_num:200000")
	pm.img.Save(pm.path)

}
