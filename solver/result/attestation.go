package result

import (
	digest "github.com/opencontainers/go-digest"
)

type Attestation interface {
	isAttestation()
}

type InTotoAttestation struct {
	PredicateType   string
	PredicateRefKey string
	PredicatePath   string
	Subjects        []InTotoSubject
}

func (a *InTotoAttestation) isAttestation() {}

type InTotoSubject interface {
	isInTotoSubject()
}

type InTotoSubjectSelf struct{}

func (as *InTotoSubjectSelf) isInTotoSubject() {}

type InTotoSubjectRaw struct {
	Name   string
	Digest []digest.Digest
}

func (as *InTotoSubjectRaw) isInTotoSubject() {}

func (as *InTotoSubjectRaw) DigestMap() map[string]string {
	m := map[string]string{}
	for _, d := range as.Digest {
		m[d.Algorithm().String()] = d.Encoded()
	}
	return m
}

func DigestToDigestMap(d digest.Digest) map[string]string {
	return map[string]string{
		d.Algorithm().String(): d.Encoded(),
	}
}