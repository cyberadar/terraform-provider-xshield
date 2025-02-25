// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AssetSynchronization struct {
	AttackSurface *bool   `json:"attackSurface,omitempty"`
	BlastRadius   *bool   `json:"blastRadius,omitempty"`
	Comment       *string `json:"comment,omitempty"`
}

func (o *AssetSynchronization) GetAttackSurface() *bool {
	if o == nil {
		return nil
	}
	return o.AttackSurface
}

func (o *AssetSynchronization) GetBlastRadius() *bool {
	if o == nil {
		return nil
	}
	return o.BlastRadius
}

func (o *AssetSynchronization) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}
