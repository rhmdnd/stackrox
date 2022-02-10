import React from 'react';
import { ExpandableSection } from '@patternfly/react-core';

import { Descriptor } from 'Containers/Policies/Wizard/Form/descriptors';
import PolicyCriteriaKey from './PolicyCriteriaKey';

type PolicyCriteriaCategoryProps = {
    category: string;
    keys: Descriptor[];
};

function PolicyCriteriaCategory({ category, keys }: PolicyCriteriaCategoryProps) {
    const [isExpanded, setIsExpanded] = React.useState(false);

    function onToggle(expanded: boolean) {
        setIsExpanded(expanded);
    }

    return (
        <ExpandableSection isExpanded={isExpanded} onToggle={onToggle} toggleText={category}>
            {keys.map((key) => (
                <PolicyCriteriaKey fieldKey={key} key={key.name} />
            ))}
        </ExpandableSection>
    );
}

export default PolicyCriteriaCategory;
