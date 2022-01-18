name: SampleCollection
type: SampleCollection
description: Collectioncontainer which holds a sample.Sample
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
    entities:
        type: sample.SampleEntity
        description: the data contains a sample.Sample
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: sample.samplecollection.entities.placeholder
            hint: ""
            label: sample.samplecollection.entities.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
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
            placeholder: sample.samplecollection.links.placeholder
            hint: ""
            label: sample.samplecollection.links.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
