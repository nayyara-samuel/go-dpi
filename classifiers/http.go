package classifiers

import (
	"github.com/google/gopacket/layers"
	"github.com/mushorg/go-dpi"
)

type HttpClassifier struct{}

func (_ HttpClassifier) HeuristicClassify(flow *godpi.Flow) bool {
	if len(flow.Packets) == 0 {
		return false
	}
	for _, packet := range flow.Packets {
		if layer := (*packet).Layer(layers.LayerTypeTCP); layer != nil {
			srcPort := layer.(*layers.TCP).SrcPort
			dstPort := layer.(*layers.TCP).DstPort
			if srcPort != 80 && dstPort != 80 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (_ HttpClassifier) GetProtocol() godpi.Protocol {
	return godpi.Http
}