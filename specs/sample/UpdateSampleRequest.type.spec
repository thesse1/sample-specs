name: UpdateSampleRequest
type: UpdateSampleRequest
description: request message for UpdateSample
lifecycle: null
__proto:
    package: sample
    targetfile: reqmsgs.proto
    imports:
        - google/protobuf/field_mask.proto
        - sample/sample.proto
    options:
        go_package: github.com/thesse1/sample-specs/dist/pb/sample;samplepb
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.example.tutorial.sample
fields:
    body:
        type: .sample.Sample
        description: Body with sample.Sample
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: sample.updatesamplerequest.body.placeholder
            hint: ""
            label: sample.updatesamplerequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    smp:
        type: string
        description: The query param smp stands for the id of a  Sample.
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: sample.updatesamplerequest.smp.placeholder
            hint: ""
            label: sample.updatesamplerequest.smp.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    update_mask:
        type: google.protobuf.FieldMask
        description: Needed to patch a record
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: sample.updatesamplerequest.updatemask.placeholder
            hint: ""
            label: sample.updatesamplerequest.updatemask.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
