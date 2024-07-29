// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-xshield-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-xshield-sdk/internal/sdk/models/shared"
)

func (r *TemplateResourceModel) ToSharedTemplate() *shared.Template {
	accessPolicyTemplate := new(bool)
	if !r.AccessPolicyTemplate.IsUnknown() && !r.AccessPolicyTemplate.IsNull() {
		*accessPolicyTemplate = r.AccessPolicyTemplate.ValueBool()
	} else {
		accessPolicyTemplate = nil
	}
	colortokensManaged := new(bool)
	if !r.ColortokensManaged.IsUnknown() && !r.ColortokensManaged.IsNull() {
		*colortokensManaged = r.ColortokensManaged.ValueBool()
	} else {
		colortokensManaged = nil
	}
	templateCategory := new(string)
	if !r.TemplateCategory.IsUnknown() && !r.TemplateCategory.IsNull() {
		*templateCategory = r.TemplateCategory.ValueString()
	} else {
		templateCategory = nil
	}
	templateDescription := new(string)
	if !r.TemplateDescription.IsUnknown() && !r.TemplateDescription.IsNull() {
		*templateDescription = r.TemplateDescription.ValueString()
	} else {
		templateDescription = nil
	}
	id := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id = r.ID.ValueString()
	} else {
		id = nil
	}
	templateName := new(string)
	if !r.TemplateName.IsUnknown() && !r.TemplateName.IsNull() {
		*templateName = r.TemplateName.ValueString()
	} else {
		templateName = nil
	}
	var templatePaths []shared.MetadataPath = []shared.MetadataPath{}
	for _, templatePathsItem := range r.TemplatePaths {
		id1 := new(string)
		if !templatePathsItem.ID.IsUnknown() && !templatePathsItem.ID.IsNull() {
			*id1 = templatePathsItem.ID.ValueString()
		} else {
			id1 = nil
		}
		destinationAssetID := new(string)
		if !templatePathsItem.DestinationAssetID.IsUnknown() && !templatePathsItem.DestinationAssetID.IsNull() {
			*destinationAssetID = templatePathsItem.DestinationAssetID.ValueString()
		} else {
			destinationAssetID = nil
		}
		var destinationNamedNetwork *shared.MetadataNamedNetworkReference
		if templatePathsItem.DestinationNamedNetwork != nil {
			namedNetworkID := new(string)
			if !templatePathsItem.DestinationNamedNetwork.NamedNetworkID.IsUnknown() && !templatePathsItem.DestinationNamedNetwork.NamedNetworkID.IsNull() {
				*namedNetworkID = templatePathsItem.DestinationNamedNetwork.NamedNetworkID.ValueString()
			} else {
				namedNetworkID = nil
			}
			namedNetworkName := new(string)
			if !templatePathsItem.DestinationNamedNetwork.NamedNetworkName.IsUnknown() && !templatePathsItem.DestinationNamedNetwork.NamedNetworkName.IsNull() {
				*namedNetworkName = templatePathsItem.DestinationNamedNetwork.NamedNetworkName.ValueString()
			} else {
				namedNetworkName = nil
			}
			destinationNamedNetwork = &shared.MetadataNamedNetworkReference{
				NamedNetworkID:   namedNetworkID,
				NamedNetworkName: namedNetworkName,
			}
		}
		var destinationTagBasedPolicy *shared.MetadataTagBasedPolicyReference
		if templatePathsItem.DestinationTagBasedPolicy != nil {
			criteria := new(string)
			if !templatePathsItem.DestinationTagBasedPolicy.Criteria.IsUnknown() && !templatePathsItem.DestinationTagBasedPolicy.Criteria.IsNull() {
				*criteria = templatePathsItem.DestinationTagBasedPolicy.Criteria.ValueString()
			} else {
				criteria = nil
			}
			criteriaAsParams := new(string)
			if !templatePathsItem.DestinationTagBasedPolicy.CriteriaAsParams.IsUnknown() && !templatePathsItem.DestinationTagBasedPolicy.CriteriaAsParams.IsNull() {
				*criteriaAsParams = templatePathsItem.DestinationTagBasedPolicy.CriteriaAsParams.ValueString()
			} else {
				criteriaAsParams = nil
			}
			tagBasedPolicyID := new(string)
			if !templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyID.IsUnknown() && !templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyID.IsNull() {
				*tagBasedPolicyID = templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyID.ValueString()
			} else {
				tagBasedPolicyID = nil
			}
			tagBasedPolicyName := new(string)
			if !templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyName.IsUnknown() && !templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyName.IsNull() {
				*tagBasedPolicyName = templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyName.ValueString()
			} else {
				tagBasedPolicyName = nil
			}
			destinationTagBasedPolicy = &shared.MetadataTagBasedPolicyReference{
				Criteria:           criteria,
				CriteriaAsParams:   criteriaAsParams,
				TagBasedPolicyID:   tagBasedPolicyID,
				TagBasedPolicyName: tagBasedPolicyName,
			}
		}
		direction := new(string)
		if !templatePathsItem.Direction.IsUnknown() && !templatePathsItem.Direction.IsNull() {
			*direction = templatePathsItem.Direction.ValueString()
		} else {
			direction = nil
		}
		domain := new(string)
		if !templatePathsItem.Domain.IsUnknown() && !templatePathsItem.Domain.IsNull() {
			*domain = templatePathsItem.Domain.ValueString()
		} else {
			domain = nil
		}
		dstIP := new(string)
		if !templatePathsItem.DstIP.IsUnknown() && !templatePathsItem.DstIP.IsNull() {
			*dstIP = templatePathsItem.DstIP.ValueString()
		} else {
			dstIP = nil
		}
		dstProcess := new(string)
		if !templatePathsItem.DstProcess.IsUnknown() && !templatePathsItem.DstProcess.IsNull() {
			*dstProcess = templatePathsItem.DstProcess.ValueString()
		} else {
			dstProcess = nil
		}
		method := new(string)
		if !templatePathsItem.Method.IsUnknown() && !templatePathsItem.Method.IsNull() {
			*method = templatePathsItem.Method.ValueString()
		} else {
			method = nil
		}
		port := new(string)
		if !templatePathsItem.Port.IsUnknown() && !templatePathsItem.Port.IsNull() {
			*port = templatePathsItem.Port.ValueString()
		} else {
			port = nil
		}
		portName := new(string)
		if !templatePathsItem.PortName.IsUnknown() && !templatePathsItem.PortName.IsNull() {
			*portName = templatePathsItem.PortName.ValueString()
		} else {
			portName = nil
		}
		protocol := new(string)
		if !templatePathsItem.Protocol.IsUnknown() && !templatePathsItem.Protocol.IsNull() {
			*protocol = templatePathsItem.Protocol.ValueString()
		} else {
			protocol = nil
		}
		sourceAssetID := new(string)
		if !templatePathsItem.SourceAssetID.IsUnknown() && !templatePathsItem.SourceAssetID.IsNull() {
			*sourceAssetID = templatePathsItem.SourceAssetID.ValueString()
		} else {
			sourceAssetID = nil
		}
		var sourceNamedNetwork *shared.MetadataNamedNetworkReference
		if templatePathsItem.SourceNamedNetwork != nil {
			namedNetworkId1 := new(string)
			if !templatePathsItem.SourceNamedNetwork.NamedNetworkID.IsUnknown() && !templatePathsItem.SourceNamedNetwork.NamedNetworkID.IsNull() {
				*namedNetworkId1 = templatePathsItem.SourceNamedNetwork.NamedNetworkID.ValueString()
			} else {
				namedNetworkId1 = nil
			}
			namedNetworkName1 := new(string)
			if !templatePathsItem.SourceNamedNetwork.NamedNetworkName.IsUnknown() && !templatePathsItem.SourceNamedNetwork.NamedNetworkName.IsNull() {
				*namedNetworkName1 = templatePathsItem.SourceNamedNetwork.NamedNetworkName.ValueString()
			} else {
				namedNetworkName1 = nil
			}
			sourceNamedNetwork = &shared.MetadataNamedNetworkReference{
				NamedNetworkID:   namedNetworkId1,
				NamedNetworkName: namedNetworkName1,
			}
		}
		var sourceTagBasedPolicy *shared.MetadataTagBasedPolicyReference
		if templatePathsItem.SourceTagBasedPolicy != nil {
			criteria1 := new(string)
			if !templatePathsItem.SourceTagBasedPolicy.Criteria.IsUnknown() && !templatePathsItem.SourceTagBasedPolicy.Criteria.IsNull() {
				*criteria1 = templatePathsItem.SourceTagBasedPolicy.Criteria.ValueString()
			} else {
				criteria1 = nil
			}
			criteriaAsParams1 := new(string)
			if !templatePathsItem.SourceTagBasedPolicy.CriteriaAsParams.IsUnknown() && !templatePathsItem.SourceTagBasedPolicy.CriteriaAsParams.IsNull() {
				*criteriaAsParams1 = templatePathsItem.SourceTagBasedPolicy.CriteriaAsParams.ValueString()
			} else {
				criteriaAsParams1 = nil
			}
			tagBasedPolicyId1 := new(string)
			if !templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyID.IsUnknown() && !templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyID.IsNull() {
				*tagBasedPolicyId1 = templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyID.ValueString()
			} else {
				tagBasedPolicyId1 = nil
			}
			tagBasedPolicyName1 := new(string)
			if !templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyName.IsUnknown() && !templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyName.IsNull() {
				*tagBasedPolicyName1 = templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyName.ValueString()
			} else {
				tagBasedPolicyName1 = nil
			}
			sourceTagBasedPolicy = &shared.MetadataTagBasedPolicyReference{
				Criteria:           criteria1,
				CriteriaAsParams:   criteriaAsParams1,
				TagBasedPolicyID:   tagBasedPolicyId1,
				TagBasedPolicyName: tagBasedPolicyName1,
			}
		}
		srcIP := new(string)
		if !templatePathsItem.SrcIP.IsUnknown() && !templatePathsItem.SrcIP.IsNull() {
			*srcIP = templatePathsItem.SrcIP.ValueString()
		} else {
			srcIP = nil
		}
		srcProcess := new(string)
		if !templatePathsItem.SrcProcess.IsUnknown() && !templatePathsItem.SrcProcess.IsNull() {
			*srcProcess = templatePathsItem.SrcProcess.ValueString()
		} else {
			srcProcess = nil
		}
		uri := new(string)
		if !templatePathsItem.URI.IsUnknown() && !templatePathsItem.URI.IsNull() {
			*uri = templatePathsItem.URI.ValueString()
		} else {
			uri = nil
		}
		templatePaths = append(templatePaths, shared.MetadataPath{
			ID:                        id1,
			DestinationAssetID:        destinationAssetID,
			DestinationNamedNetwork:   destinationNamedNetwork,
			DestinationTagBasedPolicy: destinationTagBasedPolicy,
			Direction:                 direction,
			Domain:                    domain,
			DstIP:                     dstIP,
			DstProcess:                dstProcess,
			Method:                    method,
			Port:                      port,
			PortName:                  portName,
			Protocol:                  protocol,
			SourceAssetID:             sourceAssetID,
			SourceNamedNetwork:        sourceNamedNetwork,
			SourceTagBasedPolicy:      sourceTagBasedPolicy,
			SrcIP:                     srcIP,
			SrcProcess:                srcProcess,
			URI:                       uri,
		})
	}
	var templatePorts []shared.MetadataPort = []shared.MetadataPort{}
	for _, templatePortsItem := range r.TemplatePorts {
		listenPort := new(string)
		if !templatePortsItem.ListenPort.IsUnknown() && !templatePortsItem.ListenPort.IsNull() {
			*listenPort = templatePortsItem.ListenPort.ValueString()
		} else {
			listenPort = nil
		}
		listenPortName := new(string)
		if !templatePortsItem.ListenPortName.IsUnknown() && !templatePortsItem.ListenPortName.IsNull() {
			*listenPortName = templatePortsItem.ListenPortName.ValueString()
		} else {
			listenPortName = nil
		}
		listenPortProtocol := new(string)
		if !templatePortsItem.ListenPortProtocol.IsUnknown() && !templatePortsItem.ListenPortProtocol.IsNull() {
			*listenPortProtocol = templatePortsItem.ListenPortProtocol.ValueString()
		} else {
			listenPortProtocol = nil
		}
		listenPortReviewed := new(shared.MetadataPortState)
		if !templatePortsItem.ListenPortReviewed.IsUnknown() && !templatePortsItem.ListenPortReviewed.IsNull() {
			*listenPortReviewed = shared.MetadataPortState(templatePortsItem.ListenPortReviewed.ValueString())
		} else {
			listenPortReviewed = nil
		}
		var listenProcessNames []string = []string{}
		for _, listenProcessNamesItem := range templatePortsItem.ListenProcessNames {
			listenProcessNames = append(listenProcessNames, listenProcessNamesItem.ValueString())
		}
		id2 := new(string)
		if !templatePortsItem.ID.IsUnknown() && !templatePortsItem.ID.IsNull() {
			*id2 = templatePortsItem.ID.ValueString()
		} else {
			id2 = nil
		}
		templatePorts = append(templatePorts, shared.MetadataPort{
			ListenPort:         listenPort,
			ListenPortName:     listenPortName,
			ListenPortProtocol: listenPortProtocol,
			ListenPortReviewed: listenPortReviewed,
			ListenProcessNames: listenProcessNames,
			ID:                 id2,
		})
	}
	templateType := new(shared.TemplateType)
	if !r.TemplateType.IsUnknown() && !r.TemplateType.IsNull() {
		*templateType = shared.TemplateType(r.TemplateType.ValueString())
	} else {
		templateType = nil
	}
	out := shared.Template{
		AccessPolicyTemplate: accessPolicyTemplate,
		ColortokensManaged:   colortokensManaged,
		TemplateCategory:     templateCategory,
		TemplateDescription:  templateDescription,
		ID:                   id,
		TemplateName:         templateName,
		TemplatePaths:        templatePaths,
		TemplatePorts:        templatePorts,
		TemplateType:         templateType,
	}
	return &out
}

func (r *TemplateResourceModel) RefreshFromSharedTemplate(resp *shared.Template) {
	if resp != nil {
		r.AccessPolicyTemplate = types.BoolPointerValue(resp.AccessPolicyTemplate)
		r.ColortokensManaged = types.BoolPointerValue(resp.ColortokensManaged)
		r.ID = types.StringPointerValue(resp.ID)
		r.TemplateCategory = types.StringPointerValue(resp.TemplateCategory)
		r.TemplateDescription = types.StringPointerValue(resp.TemplateDescription)
		r.TemplateName = types.StringPointerValue(resp.TemplateName)
		r.TemplatePaths = []tfTypes.MetadataPath{}
		if len(r.TemplatePaths) > len(resp.TemplatePaths) {
			r.TemplatePaths = r.TemplatePaths[:len(resp.TemplatePaths)]
		}
		for templatePathsCount, templatePathsItem := range resp.TemplatePaths {
			var templatePaths1 tfTypes.MetadataPath
			templatePaths1.DestinationAssetID = types.StringPointerValue(templatePathsItem.DestinationAssetID)
			if templatePathsItem.DestinationNamedNetwork == nil {
				templatePaths1.DestinationNamedNetwork = nil
			} else {
				templatePaths1.DestinationNamedNetwork = &tfTypes.MetadataNamedNetworkReference{}
				templatePaths1.DestinationNamedNetwork.NamedNetworkID = types.StringPointerValue(templatePathsItem.DestinationNamedNetwork.NamedNetworkID)
				templatePaths1.DestinationNamedNetwork.NamedNetworkName = types.StringPointerValue(templatePathsItem.DestinationNamedNetwork.NamedNetworkName)
			}
			if templatePathsItem.DestinationTagBasedPolicy == nil {
				templatePaths1.DestinationTagBasedPolicy = nil
			} else {
				templatePaths1.DestinationTagBasedPolicy = &tfTypes.MetadataTagBasedPolicyReference{}
				templatePaths1.DestinationTagBasedPolicy.Criteria = types.StringPointerValue(templatePathsItem.DestinationTagBasedPolicy.Criteria)
				templatePaths1.DestinationTagBasedPolicy.CriteriaAsParams = types.StringPointerValue(templatePathsItem.DestinationTagBasedPolicy.CriteriaAsParams)
				templatePaths1.DestinationTagBasedPolicy.TagBasedPolicyID = types.StringPointerValue(templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyID)
				templatePaths1.DestinationTagBasedPolicy.TagBasedPolicyName = types.StringPointerValue(templatePathsItem.DestinationTagBasedPolicy.TagBasedPolicyName)
			}
			templatePaths1.Direction = types.StringPointerValue(templatePathsItem.Direction)
			templatePaths1.Domain = types.StringPointerValue(templatePathsItem.Domain)
			templatePaths1.DstIP = types.StringPointerValue(templatePathsItem.DstIP)
			templatePaths1.DstProcess = types.StringPointerValue(templatePathsItem.DstProcess)
			templatePaths1.ID = types.StringPointerValue(templatePathsItem.ID)
			templatePaths1.Method = types.StringPointerValue(templatePathsItem.Method)
			templatePaths1.Port = types.StringPointerValue(templatePathsItem.Port)
			templatePaths1.PortName = types.StringPointerValue(templatePathsItem.PortName)
			templatePaths1.Protocol = types.StringPointerValue(templatePathsItem.Protocol)
			templatePaths1.SourceAssetID = types.StringPointerValue(templatePathsItem.SourceAssetID)
			if templatePathsItem.SourceNamedNetwork == nil {
				templatePaths1.SourceNamedNetwork = nil
			} else {
				templatePaths1.SourceNamedNetwork = &tfTypes.MetadataNamedNetworkReference{}
				templatePaths1.SourceNamedNetwork.NamedNetworkID = types.StringPointerValue(templatePathsItem.SourceNamedNetwork.NamedNetworkID)
				templatePaths1.SourceNamedNetwork.NamedNetworkName = types.StringPointerValue(templatePathsItem.SourceNamedNetwork.NamedNetworkName)
			}
			if templatePathsItem.SourceTagBasedPolicy == nil {
				templatePaths1.SourceTagBasedPolicy = nil
			} else {
				templatePaths1.SourceTagBasedPolicy = &tfTypes.MetadataTagBasedPolicyReference{}
				templatePaths1.SourceTagBasedPolicy.Criteria = types.StringPointerValue(templatePathsItem.SourceTagBasedPolicy.Criteria)
				templatePaths1.SourceTagBasedPolicy.CriteriaAsParams = types.StringPointerValue(templatePathsItem.SourceTagBasedPolicy.CriteriaAsParams)
				templatePaths1.SourceTagBasedPolicy.TagBasedPolicyID = types.StringPointerValue(templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyID)
				templatePaths1.SourceTagBasedPolicy.TagBasedPolicyName = types.StringPointerValue(templatePathsItem.SourceTagBasedPolicy.TagBasedPolicyName)
			}
			templatePaths1.SrcIP = types.StringPointerValue(templatePathsItem.SrcIP)
			templatePaths1.SrcProcess = types.StringPointerValue(templatePathsItem.SrcProcess)
			templatePaths1.URI = types.StringPointerValue(templatePathsItem.URI)
			if templatePathsCount+1 > len(r.TemplatePaths) {
				r.TemplatePaths = append(r.TemplatePaths, templatePaths1)
			} else {
				r.TemplatePaths[templatePathsCount].DestinationAssetID = templatePaths1.DestinationAssetID
				r.TemplatePaths[templatePathsCount].DestinationNamedNetwork = templatePaths1.DestinationNamedNetwork
				r.TemplatePaths[templatePathsCount].DestinationTagBasedPolicy = templatePaths1.DestinationTagBasedPolicy
				r.TemplatePaths[templatePathsCount].Direction = templatePaths1.Direction
				r.TemplatePaths[templatePathsCount].Domain = templatePaths1.Domain
				r.TemplatePaths[templatePathsCount].DstIP = templatePaths1.DstIP
				r.TemplatePaths[templatePathsCount].DstProcess = templatePaths1.DstProcess
				r.TemplatePaths[templatePathsCount].ID = templatePaths1.ID
				r.TemplatePaths[templatePathsCount].Method = templatePaths1.Method
				r.TemplatePaths[templatePathsCount].Port = templatePaths1.Port
				r.TemplatePaths[templatePathsCount].PortName = templatePaths1.PortName
				r.TemplatePaths[templatePathsCount].Protocol = templatePaths1.Protocol
				r.TemplatePaths[templatePathsCount].SourceAssetID = templatePaths1.SourceAssetID
				r.TemplatePaths[templatePathsCount].SourceNamedNetwork = templatePaths1.SourceNamedNetwork
				r.TemplatePaths[templatePathsCount].SourceTagBasedPolicy = templatePaths1.SourceTagBasedPolicy
				r.TemplatePaths[templatePathsCount].SrcIP = templatePaths1.SrcIP
				r.TemplatePaths[templatePathsCount].SrcProcess = templatePaths1.SrcProcess
				r.TemplatePaths[templatePathsCount].URI = templatePaths1.URI
			}
		}
		r.TemplatePorts = []tfTypes.MetadataPort{}
		if len(r.TemplatePorts) > len(resp.TemplatePorts) {
			r.TemplatePorts = r.TemplatePorts[:len(resp.TemplatePorts)]
		}
		for templatePortsCount, templatePortsItem := range resp.TemplatePorts {
			var templatePorts1 tfTypes.MetadataPort
			templatePorts1.ID = types.StringPointerValue(templatePortsItem.ID)
			templatePorts1.ListenPort = types.StringPointerValue(templatePortsItem.ListenPort)
			templatePorts1.ListenPortName = types.StringPointerValue(templatePortsItem.ListenPortName)
			templatePorts1.ListenPortProtocol = types.StringPointerValue(templatePortsItem.ListenPortProtocol)
			if templatePortsItem.ListenPortReviewed != nil {
				templatePorts1.ListenPortReviewed = types.StringValue(string(*templatePortsItem.ListenPortReviewed))
			} else {
				templatePorts1.ListenPortReviewed = types.StringNull()
			}
			templatePorts1.ListenProcessNames = []types.String{}
			for _, v := range templatePortsItem.ListenProcessNames {
				templatePorts1.ListenProcessNames = append(templatePorts1.ListenProcessNames, types.StringValue(v))
			}
			if templatePortsCount+1 > len(r.TemplatePorts) {
				r.TemplatePorts = append(r.TemplatePorts, templatePorts1)
			} else {
				r.TemplatePorts[templatePortsCount].ID = templatePorts1.ID
				r.TemplatePorts[templatePortsCount].ListenPort = templatePorts1.ListenPort
				r.TemplatePorts[templatePortsCount].ListenPortName = templatePorts1.ListenPortName
				r.TemplatePorts[templatePortsCount].ListenPortProtocol = templatePorts1.ListenPortProtocol
				r.TemplatePorts[templatePortsCount].ListenPortReviewed = templatePorts1.ListenPortReviewed
				r.TemplatePorts[templatePortsCount].ListenProcessNames = templatePorts1.ListenProcessNames
			}
		}
		if resp.TemplateType != nil {
			r.TemplateType = types.StringValue(string(*resp.TemplateType))
		} else {
			r.TemplateType = types.StringNull()
		}
	}
}
