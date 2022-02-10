import React, { useState, useEffect, ReactElement } from 'react';
import { FormGroup, Select, SelectOption, SelectVariant } from '@patternfly/react-core';
import { useField } from 'formik';

import { getPolicyCategories } from 'services/PoliciesService';

function PolicyCategoriesSelectField(): ReactElement {
    const [policyCategories, setPolicyCategories] = useState<string[]>([]);
    // manage state for Categories select below
    const [isCategoriesOpen, setIsCategoriesOpen] = useState(false);

    const [field, , helpers] = useField('categories');

    function onCategoriesToggle(isOpen) {
        setIsCategoriesOpen(isOpen);
    }

    function onCreateCategory(newCategory) {
        setPolicyCategories([...policyCategories, newCategory]);
    }

    function onSelectHandler(selectedCategories) {
        return (event, selection) => {
            const newSelectedCategories = selectedCategories.includes(selection)
                ? selectedCategories.filter((item) => item !== selection)
                : [...selectedCategories, selection];
            helpers.setValue(newSelectedCategories);
        };
    }

    function clearSelection() {
        setIsCategoriesOpen(false);
        helpers.setValue([]);
    }

    useEffect(() => {
        getPolicyCategories()
            .then((response) => {
                setPolicyCategories(response);
            })
            .catch(() => {});

        return () => {
            setPolicyCategories([]);
        };
    }, []);

    // So wizard footer does not overlap select menu, display 3 categories with scrollbar.
    const maxHeight = '377%'; // (3 * 40px + 2 * 8px) / 36px

    return (
        <FormGroup fieldId="policy-categories" label="Categories" isRequired>
            <Select
                variant={SelectVariant.typeaheadMulti}
                name={field.name}
                value={field.value}
                isOpen={isCategoriesOpen}
                selections={field.value}
                onSelect={onSelectHandler(field.value)}
                onToggle={onCategoriesToggle}
                onCreateOption={onCreateCategory}
                onClear={clearSelection}
                isCreatable
                maxHeight={maxHeight}
            >
                {policyCategories.map((category) => (
                    <SelectOption key={category} value={category} />
                ))}
            </Select>
        </FormGroup>
    );
}

export default PolicyCategoriesSelectField;
