name: SampleEntity
type: SampleEntity
description: SimpleEntity
lifecycle: null
__proto:
    package: sample
    targetfile: sample.proto
    imports:
        - furo/furo.proto
    options:
        go_package: github.com/thesse1/sample-specs/dist/pb/sample;samplepb
        java_multiple_files: "true"
        java_outer_classname: SampleProto
        java_package: com.example.tutorial.sample
fields:
    data:
        type: sample.Sample
        description: the data contains a sample.Sample
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: sample.sampleentity.data.placeholder
            hint: ""
            label: sample.sampleentity.data.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    links:
        type: furo.Link
        description: the Hateoas links
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: sample.sampleentity.links.placeholder
            hint: ""
            label: sample.sampleentity.links.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    meta:
        type: furo.Meta
        description: Meta for the response
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: sample.sampleentity.meta.placeholder
            hint: ""
            label: sample.sampleentity.meta.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
