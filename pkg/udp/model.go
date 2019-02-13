package udp

type Event int

const (
	// Receive
	EventCollisionWithCar Event = 10
	EventCollisionWithEnv Event = 11
	EventNewSession       Event = 50
	EventNewConnection    Event = 51
	EventConnectionClosed Event = 52
	EventCarUpdate        Event = 53
	EventCarInfo          Event = 54
	EventEndSession       Event = 55
	EventVersion          Event = 56
	EventChat             Event = 57
	EventClientLoaded     Event = 58
	EventSessionInfo      Event = 59
	EventError            Event = 60
	EventLapCompleted     Event = 73
	EventClientEvent      Event = 130
	EventSectorCompleted  Event = 88

	// Send
	EventRealtimeposInterval Event = 200
	EventGetCarInfo          Event = 201
	EventSendChat            Event = 202
	EventBroadcastChat       Event = 203
	EventGetSessionInfo      Event = 204
	EventSetSessionInfo      Event = 205
	EventKickUser            Event = 206
	EventNextSession         Event = 207
	EventRestartSession      Event = 208
	EventAdminCommand        Event = 209
)

type Message interface {
	Event() Event
}

type ServerError struct {
	error
}

func (ServerError) Event() Event {
	return EventError
}

type CarID uint8

type LapCompletedInternal struct {
	CarID     CarID
	LapTime   uint32
	Cuts      uint8
	CarsCount uint8
}

func (LapCompleted) Event() Event {
	return EventLapCompleted
}

type LapCompleted struct {
	LapCompletedInternal

	Cars []*LapCompletedCar
}

type LapCompletedCar struct {
	CarID     CarID
	LapTime   uint32
	Laps      uint16
	Completed uint8
}

type Vec struct {
	X, Y, Z float32
}

type CollisionWithCar struct {
	CarID       CarID
	OtherCarID  uint8
	ImpactSpeed float32
	WorldPos    Vec
	RelPos      Vec
}

func (CollisionWithCar) Event() Event {
	return EventCollisionWithCar
}

type CollisionWithEnvironment struct {
	CarID       CarID
	ImpactSpeed float32
	WorldPos    Vec
	RelPos      Vec
}

func (CollisionWithEnvironment) Event() Event {
	return EventCollisionWithEnv
}

type SessionCarInfo struct {
	CarID      CarID
	DriverName string
	DriverGUID string
	CarMode    string
	CarSkin    string

	event Event
}

func (s SessionCarInfo) Event() Event {
	return s.event
}

type Chat struct {
	CarID   CarID
	Message string
}

func (Chat) Event() Event {
	return EventChat
}

type CarInfo struct {
	CarID       CarID
	IsConnected bool
	CarModel    string
	CarSkin     string
	DriverName  string
	DriverTeam  string
	DriverGUID  string
}

func (CarInfo) Event() Event {
	return EventCarInfo
}

type CarUpdate struct {
	CarID               CarID
	Pos                 Vec
	Velocity            Vec
	Gear                uint8
	EngineRPM           uint16
	NormalisedSplinePos float32
}

func (CarUpdate) Event() Event {
	return EventCarUpdate
}

type EndSession string

func (EndSession) Event() Event {
	return EventEndSession
}

type Version uint8

func (Version) Event() Event {
	return EventVersion
}

type ClientLoaded CarID

func (ClientLoaded) Event() Event {
	return EventClientLoaded
}

type SessionInfo struct {
	Version             uint8
	SessionIndex        uint8
	CurrentSessionIndex uint8
	SessionCount        uint8
	ServerName          string
	Track               string
	TrackConfig         string
	Name                string
	Type                uint8
	Time                uint16
	Laps                uint16
	WaitTime            uint16
	AmbientTemp         uint8
	RoadTemp            uint8
	WeatherGraphics     string
	ElapsedMilliseconds int32

	event Event
}

func (s SessionInfo) Event() Event {
	return s.event
}

type SectorCompleted struct {
	Num uint8
	Anothernum uint8
	Time uint16
}