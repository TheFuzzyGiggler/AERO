package main

import (
	"encoding/json"
	"time"
)

type EDDN struct {
	SchemaRef string          `json:"$schemaRef"`
	Header    EDDNHeader      `json:"header"`
	Message   json.RawMessage `json:"message"`
}

type EDDNHeader struct {
	UploaderID       string    `json:"uploaderID"`
	SoftwareName     string    `json:"softwareName"`
	SoftwareVersion  string    `json:"softwareVersion"`
	GameVersion      string    `json:"gameversion,omitempty"`
	GameBuild        string    `json:"gamebuild,omitempty"`
	GatewayTimestamp time.Time `json:"gatewayTimestamp"`
}

type ApproachSettlementMessage struct {
	Timestamp         string           `json:"timestamp"`
	Event             string           `json:"event"`
	StarSystem        string           `json:"StarSystem"`
	StarPos           [3]float64       `json:"StarPos"`
	SystemAddress     int64            `json:"SystemAddress"`
	Name              string           `json:"Name"`
	MarketID          int64            `json:"MarketID,omitempty"`
	BodyID            int              `json:"BodyID"`
	BodyName          string           `json:"BodyName"`
	Latitude          float64          `json:"Latitude"`
	Longitude         float64          `json:"Longitude"`
	Horizons          bool             `json:"horizons,omitempty"`
	Odyssey           bool             `json:"odyssey,omitempty"`
	StationGovernment *string          `json:"StationGovernment,omitempty"`
	StationAllegiance *string          `json:"StationAllegiance,omitempty"`
	StationEconomies  []StationEconomy `json:"StationEconomies,omitempty"`
	StationFaction    *StationFaction  `json:"StationFaction,omitempty"`
	StationServices   []string         `json:"StationServices,omitempty"`
	StationEconomy    *string          `json:"StationEconomy,omitempty"`
}

type StationEconomy struct {
	Name       string  `json:"Name"`
	Proportion float64 `json:"Proportion"`
}

type StationFaction struct {
	Name         string `json:"Name"`
	FactionState string `json:"FactionState"`
}

type CommodityMessage struct {
	SystemName           string           `json:"systemName"`
	StationName          string           `json:"stationName"`
	StationType          string           `json:"stationType,omitempty"`
	CarrierDockingAccess string           `json:"carrierDockingAccess,omitempty"`
	MarketID             float64          `json:"marketId"` // Changed to float64
	Horizons             bool             `json:"horizons,omitempty"`
	Odyssey              bool             `json:"odyssey,omitempty"`
	Timestamp            string           `json:"timestamp"`
	Commodities          []CommodityEntry `json:"commodities"`
	Economies            []Economy        `json:"economies,omitempty"`
	Prohibited           []string         `json:"prohibited,omitempty"`
}

type CommodityEntry struct {
	Name          string  `json:"name"`
	MeanPrice     float64 `json:"meanPrice,omitempty"`     // Changed to float64
	BuyPrice      float64 `json:"buyPrice,omitempty"`      // Changed to float64
	Stock         float64 `json:"stock,omitempty"`         // Changed to float64
	StockBracket  float64 `json:"stockBracket,omitempty"`  // Changed to float64
	SellPrice     float64 `json:"sellPrice,omitempty"`     // Changed to float64
	Demand        float64 `json:"demand,omitempty"`        // Changed to float64
	DemandBracket float64 `json:"demandBracket,omitempty"` // Changed to float64
}

type Economy struct {
	Name       string  `json:"name"`
	Proportion float64 `json:"proportion"`
}

// Struct for FCMaterials messages
type FCMaterialsMessage struct {
	Timestamp string `json:"timestamp"`
	Event     string `json:"event"`
	MarketID  int64  `json:"MarketID"`
	CarrierID string `json:"CarrierID"`
	Items     Items  `json:"Items"`
}

type Items struct {
	Purchases []Purchase `json:"purchases"`
}

type Purchase struct {
	Name        string `json:"name"`
	Outstanding int    `json:"outstanding"`
	Price       int    `json:"price"`
	Total       int    `json:"total"`
}

type BlackMarketMessage struct {
	SystemName   string `json:"systemName"`
	StationName  string `json:"stationName"`
	MarketID     int64  `json:"marketId,omitempty"` // Renamed from MarketID
	Timestamp    string `json:"timestamp"`
	Type         string `json:"name"` // Renamed from Type
	SellPrice    int    `json:"sellPrice"`
	IllegalGoods bool   `json:"prohibited"` // Renamed from IllegalGoods
}

type JournalMessage struct {
	Timestamp             string     `json:"timestamp"`
	Event                 string     `json:"event"`
	StarSystem            string     `json:"StarSystem"`
	StarPos               [3]float64 `json:"StarPos"`
	BodyName              string     `json:"BodyName,omitempty"`
	BodyType              string     `json:"BodyType,omitempty"`
	DistanceFromArrivalLS float64    `json:"DistanceFromArrivalLS,omitempty"`
	System                string     `json:"System,omitempty"`
	WasDiscovered         bool       `json:"WasDiscovered,omitempty"`
	WasMapped             bool       `json:"WasMapped,omitempty"`
	SystemAddress         float64    `json:"SystemAddress"`
	Horizons              bool       `json:"horizons,omitempty"`
	Odyssey               bool       `json:"odyssey,omitempty"`
	Factions              []Faction  `json:"Factions,omitempty"`
	BodyID                float64    `json:"BodyID,omitempty"`
	ScanType              string     `json:"ScanType,omitempty"`
	Population            float64    `json:"Population,omitempty"`
	PowerplayState        string     `json:"PowerplayState,omitempty"`
	SystemEconomy         string     `json:"SystemEconomy,omitempty"`
	SystemSecondEconomy   string     `json:"SystemSecondEconomy,omitempty"`
	SystemAllegiance      string     `json:"SystemAllegiance,omitempty"`
	SystemSecurity        string     `json:"SystemSecurity,omitempty"`
	Signals               []Signal   `json:"Signals,omitempty"`

	Taxi bool `json:"taxi,omitempty"`
}

type Powers struct {
	Power string `json:"Power"`
}
type Faction struct {
	Name             string  `json:"Name"`
	Influence        float64 `json:"Influence"`
	FactionState     string  `json:"FactionState"`
	Allegiance       string  `json:"Allegiance"`
	Government       string  `json:"Government"`
	Happiness        string  `json:"Happiness"`
	ActiveStates     []State `json:"ActiveStates,omitempty"`
	PendingStates    []State `json:"PendingStates,omitempty"`
	RecoveringStates []State `json:"RecoveringStates,omitempty"`
}

type State struct {
	State string `json:"State"`
}

type FCMaterialsJournalMessage struct {
	Timestamp   string        `json:"timestamp"`
	Event       string        `json:"event"`
	MarketID    int64         `json:"MarketID"`
	CarrierName string        `json:"CarrierName"`
	CarrierID   string        `json:"CarrierID"`
	Items       []JournalItem `json:"Items"`
	Horizons    bool          `json:"horizons,omitempty"`
	Odyssey     bool          `json:"odyssey,omitempty"`
}

type JournalItem struct {
	ID     int    `json:"id"`
	Name   string `json:"Name"`
	Price  int    `json:"Price"`
	Stock  int    `json:"Stock"`
	Demand int    `json:"Demand"`
}

type OutfittingMessage struct {
	SystemName  string   `json:"systemName"`  // Renamed to "StarSystem" in schema
	StationName string   `json:"stationName"` // Kept as "StationName"
	MarketID    int64    `json:"marketId"`    // Renamed to "MarketID"
	Horizons    bool     `json:"horizons,omitempty"`
	Odyssey     bool     `json:"odyssey,omitempty"`
	Timestamp   string   `json:"timestamp"`
	Modules     []string `json:"modules"` // Renamed to "Items" in schema
}

type NavRouteMessage struct {
	Timestamp string     `json:"timestamp"`
	Event     string     `json:"event"`
	Horizons  bool       `json:"horizons,omitempty"`
	Odyssey   bool       `json:"odyssey,omitempty"`
	Route     []NavRoute `json:"Route"`
}

type NavRoute struct {
	StarSystem    string     `json:"StarSystem"`
	SystemAddress int64      `json:"SystemAddress"`
	StarPos       [3]float64 `json:"StarPos"`
	StarClass     string     `json:"StarClass"`
}

type FSSSignalDiscoveredMessage struct {
	Event         string           `json:"event"`
	Timestamp     string           `json:"timestamp"`
	SystemAddress int64            `json:"SystemAddress"`
	StarSystem    string           `json:"StarSystem"`
	StarPos       [3]float64       `json:"StarPos"`
	Horizons      bool             `json:"horizons,omitempty"`
	Odyssey       bool             `json:"odyssey,omitempty"`
	Signals       []FSSSignalEvent `json:"signals"`
}

type FSSSignalEvent struct {
	Timestamp       string `json:"timestamp"`
	SignalName      string `json:"SignalName"`
	SignalType      string `json:"SignalType,omitempty"`
	IsStation       bool   `json:"IsStation,omitempty"`
	USSType         string `json:"USSType,omitempty"`
	SpawningState   string `json:"SpawningState,omitempty"`
	SpawningFaction string `json:"SpawningFaction,omitempty"`
	ThreatLevel     int    `json:"ThreatLevel,omitempty"`
}

type FSSAllBodiesFoundMessage struct {
	Timestamp     string     `json:"timestamp"`
	Event         string     `json:"event"`
	SystemName    string     `json:"SystemName"`
	StarPos       [3]float64 `json:"StarPos"`
	SystemAddress int64      `json:"SystemAddress"`
	Count         int        `json:"Count"`
	Horizons      bool       `json:"horizons,omitempty"`
	Odyssey       bool       `json:"odyssey,omitempty"`
}

type ScanBaryCentreMessage struct {
	Timestamp          string     `json:"timestamp"`
	Event              string     `json:"event"`
	StarSystem         string     `json:"StarSystem"`
	StarPos            [3]float64 `json:"StarPos"`
	SystemAddress      int64      `json:"SystemAddress"`
	BodyID             int        `json:"BodyID"`
	SemiMajorAxis      float64    `json:"SemiMajorAxis,omitempty"`
	Eccentricity       float64    `json:"Eccentricity,omitempty"`
	OrbitalInclination float64    `json:"OrbitalInclination,omitempty"`
	Periapsis          float64    `json:"Periapsis,omitempty"`
	OrbitalPeriod      float64    `json:"OrbitalPeriod,omitempty"`
	AscendingNode      float64    `json:"AscendingNode,omitempty"`
	MeanAnomaly        float64    `json:"MeanAnomaly,omitempty"`
	Horizons           bool       `json:"horizons,omitempty"`
	Odyssey            bool       `json:"odyssey,omitempty"`
}

type DockingDeniedMessage struct {
	Timestamp   string `json:"timestamp"`
	Event       string `json:"event"`
	MarketID    int64  `json:"MarketID"`
	StationName string `json:"StationName"`
	StationType string `json:"StationType,omitempty"`
	Reason      string `json:"Reason"`
	Horizons    bool   `json:"horizons,omitempty"`
	Odyssey     bool   `json:"odyssey,omitempty"`
}
type DockingGrantedMessage struct {
	Timestamp   string `json:"timestamp"`
	Event       string `json:"event"`
	MarketID    int64  `json:"MarketID"`
	StationName string `json:"StationName"`
	StationType string `json:"StationType,omitempty"`
	LandingPad  int    `json:"LandingPad"`
	Horizons    bool   `json:"horizons,omitempty"`
	Odyssey     bool   `json:"odyssey,omitempty"`
}
type FSSDiscoveryScanMessage struct {
	Timestamp     string     `json:"timestamp"`
	Event         string     `json:"event"`
	SystemName    string     `json:"SystemName"`
	StarPos       [3]float64 `json:"StarPos"`
	SystemAddress int64      `json:"SystemAddress"`
	BodyCount     int        `json:"BodyCount"`
	NonBodyCount  int        `json:"NonBodyCount"`
	Horizons      bool       `json:"horizons,omitempty"`
	Odyssey       bool       `json:"odyssey,omitempty"`
}
type CodexEntryMessage struct {
	Timestamp          string     `json:"timestamp"`
	Event              string     `json:"event"`
	System             string     `json:"System"`
	StarPos            [3]float64 `json:"StarPos"`
	SystemAddress      int64      `json:"SystemAddress"`
	EntryID            int        `json:"EntryID"`
	Name               string     `json:"Name"`
	Region             string     `json:"Region"`
	Category           string     `json:"Category"`
	SubCategory        string     `json:"SubCategory"`
	NearestDestination string     `json:"NearestDestination,omitempty"`
	VoucherAmount      int        `json:"VoucherAmount,omitempty"`
	Traits             []string   `json:"Traits,omitempty"`
	BodyID             int        `json:"BodyID,omitempty"`
	BodyName           string     `json:"BodyName,omitempty"`
	Latitude           float64    `json:"Latitude,omitempty"`
	Longitude          float64    `json:"Longitude,omitempty"`
	Horizons           bool       `json:"horizons,omitempty"`
	Odyssey            bool       `json:"odyssey,omitempty"`
}

type ShipyardMessage struct {
	SystemName     string   `json:"systemName"`  // Renamed to "StarSystem"
	StationName    string   `json:"stationName"` // Renamed to "StationName"
	MarketID       int64    `json:"marketId"`    // Renamed to "MarketID"
	Timestamp      string   `json:"timestamp"`
	Horizons       bool     `json:"horizons,omitempty"`
	Odyssey        bool     `json:"odyssey,omitempty"`
	AllowCobraMkIV bool     `json:"allowCobraMkIV"`
	Ships          []string `json:"ships"` // Renamed to "PriceList"
}

type FSSBodySignalsMessage struct {
	Timestamp     string     `json:"timestamp"`
	Event         string     `json:"event"`
	StarSystem    string     `json:"StarSystem"`
	StarPos       [3]float64 `json:"StarPos"`
	SystemAddress int64      `json:"SystemAddress"`
	BodyID        int        `json:"BodyID"`
	BodyName      string     `json:"BodyName"`
	Signals       []Signal   `json:"Signals"`
	Horizons      bool       `json:"horizons,omitempty"`
	Odyssey       bool       `json:"odyssey,omitempty"`
}

type Signal struct {
	Type  string `json:"Type"`
	Count int    `json:"Count"`
}

type NavBeaconScanMessage struct {
	Timestamp     string     `json:"timestamp"`
	Event         string     `json:"event"`
	StarSystem    string     `json:"StarSystem"`
	StarPos       [3]float64 `json:"StarPos"`
	SystemAddress int64      `json:"SystemAddress"`
	NumBodies     int        `json:"NumBodies"`
	Horizons      bool       `json:"horizons,omitempty"`
	Odyssey       bool       `json:"odyssey,omitempty"`
}
