package media



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var MediaItem1Interface = "org.bluez.MediaItem1"


// NewMediaItem1 create a new instance of MediaItem1
//
// Args:
// - servicePath: unique name
// - objectPath: freely definable
func NewMediaItem1(servicePath string, objectPath dbus.ObjectPath) (*MediaItem1, error) {
	a := new(MediaItem1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  servicePath,
			Iface: MediaItem1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaItem1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}

// NewMediaItem1Controller create a new instance of MediaItem1
//
// Args:
// - objectPath: [variable	prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX/playerX/itemX
func NewMediaItem1Controller(objectPath dbus.ObjectPath) (*MediaItem1, error) {
	a := new(MediaItem1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: MediaItem1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaItem1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
MediaItem1 MediaItem1 hierarchy

*/
type MediaItem1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*MediaItem1Properties
}

// MediaItem1Properties contains the exposed properties of an interface
type MediaItem1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Genre Item genre name

					Available if property Type is "audio"
					or "video"
	*/
	Genre string

	/*
	NumberOfTracks Item album number of tracks in total

					Available if property Type is "audio"
					or "video"
	*/
	NumberOfTracks uint32

	/*
	Number Item album number

					Available if property Type is "audio"
					or "video"
	*/
	Number uint32

	/*
	Duration Item duration in milliseconds

					Available if property Type is "audio"
					or "video"
	*/
	Duration uint32

	/*
	Name Item displayable name
	*/
	Name string

	/*
	Type Item type

			Possible values: "video", "audio", "folder"
	*/
	Type string

	/*
	FolderType Folder type.

			Possible values: "mixed", "titles", "albums", "artists"

			Available if property Type is "Folder"
	*/
	FolderType string

	/*
	Metadata Item metadata.

			Possible values:
	*/
	Metadata map[string]interface{}

	/*
	Album Item album name

					Available if property Type is "audio"
					or "video"
	*/
	Album string

	/*
	Player Player object path the item belongs to
	*/
	Player dbus.ObjectPath

	/*
	Playable Indicates if the item can be played

			Available if property Type is "folder"
	*/
	Playable bool

	/*
	Title Item title name

					Available if property Type is "audio"
					or "video"
	*/
	Title string

	/*
	Artist Item artist name

					Available if property Type is "audio"
					or "video"
	*/
	Artist string

}

//Lock access to properties
func (p *MediaItem1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *MediaItem1Properties) Unlock() {
	p.lock.Unlock()
}




// SetGenre set Genre value
func (a *MediaItem1) SetGenre(v string) error {
	return a.SetProperty("Genre", v)
}



// GetGenre get Genre value
func (a *MediaItem1) GetGenre() (string, error) {
	v, err := a.GetProperty("Genre")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}




// SetNumberOfTracks set NumberOfTracks value
func (a *MediaItem1) SetNumberOfTracks(v uint32) error {
	return a.SetProperty("NumberOfTracks", v)
}



// GetNumberOfTracks get NumberOfTracks value
func (a *MediaItem1) GetNumberOfTracks() (uint32, error) {
	v, err := a.GetProperty("NumberOfTracks")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}




// SetNumber set Number value
func (a *MediaItem1) SetNumber(v uint32) error {
	return a.SetProperty("Number", v)
}



// GetNumber get Number value
func (a *MediaItem1) GetNumber() (uint32, error) {
	v, err := a.GetProperty("Number")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}




// SetDuration set Duration value
func (a *MediaItem1) SetDuration(v uint32) error {
	return a.SetProperty("Duration", v)
}



// GetDuration get Duration value
func (a *MediaItem1) GetDuration() (uint32, error) {
	v, err := a.GetProperty("Duration")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}






// GetName get Name value
func (a *MediaItem1) GetName() (string, error) {
	v, err := a.GetProperty("Name")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}






// GetType get Type value
func (a *MediaItem1) GetType() (string, error) {
	v, err := a.GetProperty("Type")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}






// GetFolderType get FolderType value
func (a *MediaItem1) GetFolderType() (string, error) {
	v, err := a.GetProperty("FolderType")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}






// GetMetadata get Metadata value
func (a *MediaItem1) GetMetadata() (map[string]interface{}, error) {
	v, err := a.GetProperty("Metadata")
	if err != nil {
		return map[string]interface{}{}, err
	}
	return v.Value().(map[string]interface{}), nil
}




// SetAlbum set Album value
func (a *MediaItem1) SetAlbum(v string) error {
	return a.SetProperty("Album", v)
}



// GetAlbum get Album value
func (a *MediaItem1) GetAlbum() (string, error) {
	v, err := a.GetProperty("Album")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}






// GetPlayer get Player value
func (a *MediaItem1) GetPlayer() (dbus.ObjectPath, error) {
	v, err := a.GetProperty("Player")
	if err != nil {
		return dbus.ObjectPath(""), err
	}
	return v.Value().(dbus.ObjectPath), nil
}






// GetPlayable get Playable value
func (a *MediaItem1) GetPlayable() (bool, error) {
	v, err := a.GetProperty("Playable")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}




// SetTitle set Title value
func (a *MediaItem1) SetTitle(v string) error {
	return a.SetProperty("Title", v)
}



// GetTitle get Title value
func (a *MediaItem1) GetTitle() (string, error) {
	v, err := a.GetProperty("Title")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}




// SetArtist set Artist value
func (a *MediaItem1) SetArtist(v string) error {
	return a.SetProperty("Artist", v)
}



// GetArtist get Artist value
func (a *MediaItem1) GetArtist() (string, error) {
	v, err := a.GetProperty("Artist")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}



// Close the connection
func (a *MediaItem1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return MediaItem1 object path
func (a *MediaItem1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return MediaItem1 dbus client
func (a *MediaItem1) Client() *bluez.Client {
	return a.client
}

// Interface return MediaItem1 interface
func (a *MediaItem1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *MediaItem1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a MediaItem1Properties to map
func (a *MediaItem1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an MediaItem1Properties
func (a *MediaItem1Properties) FromMap(props map[string]interface{}) (*MediaItem1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an MediaItem1Properties
func (a *MediaItem1Properties) FromDBusMap(props map[string]dbus.Variant) (*MediaItem1Properties, error) {
	s := new(MediaItem1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *MediaItem1) ToProps() bluez.Properties {
	return a.Properties
}

// GetProperties load all available properties
func (a *MediaItem1) GetProperties() (*MediaItem1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *MediaItem1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *MediaItem1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *MediaItem1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *MediaItem1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *MediaItem1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *MediaItem1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
Play 
			Play item

			Possible Errors: org.bluez.Error.NotSupported
					 org.bluez.Error.Failed


*/
func (a *MediaItem1) Play() error {
	
	return a.client.Call("Play", 0, ).Store()
	
}

/*
AddtoNowPlaying 
			Add item to now playing list

			Possible Errors: org.bluez.Error.NotSupported
					 org.bluez.Error.Failed


*/
func (a *MediaItem1) AddtoNowPlaying() error {
	
	return a.client.Call("AddtoNowPlaying", 0, ).Store()
	
}

