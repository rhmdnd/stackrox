// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package central

func (m *CancelPullTelemetryDataRequest) Size() int                              { return m.SizeVT() }
func (m *CancelPullTelemetryDataRequest) Clone() *CancelPullTelemetryDataRequest { return m.CloneVT() }
func (m *CancelPullTelemetryDataRequest) Marshal() ([]byte, error)               { return m.MarshalVT() }
func (m *CancelPullTelemetryDataRequest) Unmarshal(dAtA []byte) error            { return m.UnmarshalVT(dAtA) }

func (m *PullTelemetryDataRequest) Size() int                        { return m.SizeVT() }
func (m *PullTelemetryDataRequest) Clone() *PullTelemetryDataRequest { return m.CloneVT() }
func (m *PullTelemetryDataRequest) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *PullTelemetryDataRequest) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *TelemetryResponsePayload) Size() int                        { return m.SizeVT() }
func (m *TelemetryResponsePayload) Clone() *TelemetryResponsePayload { return m.CloneVT() }
func (m *TelemetryResponsePayload) Marshal() ([]byte, error)         { return m.MarshalVT() }
func (m *TelemetryResponsePayload) Unmarshal(dAtA []byte) error      { return m.UnmarshalVT(dAtA) }

func (m *TelemetryResponsePayload_EndOfStream) Size() int { return m.SizeVT() }
func (m *TelemetryResponsePayload_EndOfStream) Clone() *TelemetryResponsePayload_EndOfStream {
	return m.CloneVT()
}
func (m *TelemetryResponsePayload_EndOfStream) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *TelemetryResponsePayload_EndOfStream) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *TelemetryResponsePayload_KubernetesInfo) Size() int { return m.SizeVT() }
func (m *TelemetryResponsePayload_KubernetesInfo) Clone() *TelemetryResponsePayload_KubernetesInfo {
	return m.CloneVT()
}
func (m *TelemetryResponsePayload_KubernetesInfo) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *TelemetryResponsePayload_KubernetesInfo) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *TelemetryResponsePayload_KubernetesInfo_File) Size() int { return m.SizeVT() }
func (m *TelemetryResponsePayload_KubernetesInfo_File) Clone() *TelemetryResponsePayload_KubernetesInfo_File {
	return m.CloneVT()
}
func (m *TelemetryResponsePayload_KubernetesInfo_File) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *TelemetryResponsePayload_KubernetesInfo_File) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *TelemetryResponsePayload_ClusterInfo) Size() int { return m.SizeVT() }
func (m *TelemetryResponsePayload_ClusterInfo) Clone() *TelemetryResponsePayload_ClusterInfo {
	return m.CloneVT()
}
func (m *TelemetryResponsePayload_ClusterInfo) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *TelemetryResponsePayload_ClusterInfo) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *PullTelemetryDataResponse) Size() int                         { return m.SizeVT() }
func (m *PullTelemetryDataResponse) Clone() *PullTelemetryDataResponse { return m.CloneVT() }
func (m *PullTelemetryDataResponse) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *PullTelemetryDataResponse) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *TelemetryConfig) Size() int                   { return m.SizeVT() }
func (m *TelemetryConfig) Clone() *TelemetryConfig     { return m.CloneVT() }
func (m *TelemetryConfig) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *TelemetryConfig) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }