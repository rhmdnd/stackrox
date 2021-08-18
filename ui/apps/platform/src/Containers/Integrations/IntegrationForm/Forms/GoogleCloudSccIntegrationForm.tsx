import React, { ReactElement } from 'react';
import { Checkbox, Form, PageSection, TextInput, TextArea } from '@patternfly/react-core';
import * as yup from 'yup';

import usePageState from 'Containers/Integrations/hooks/usePageState';
import useIntegrationForm from '../useIntegrationForm';
import { IntegrationFormProps } from '../integrationFormTypes';

import IntegrationFormActions from '../IntegrationFormActions';
import FormCancelButton from '../FormCancelButton';
import FormTestButton from '../FormTestButton';
import FormSaveButton from '../FormSaveButton';
import FormMessage from '../FormMessage';
import FormLabelGroup from '../FormLabelGroup';

export type GoogleCloudSccIntegration = {
    id?: string;
    name: string;
    cscc: {
        serviceAccount: string;
        sourceId: string;
    };
    uiEndpoint: string;
    type: 'cscc';
    enabled: boolean;
};

export type GoogleCloudSccIntegrationFormValues = {
    notifier: GoogleCloudSccIntegration;
    updatePassword: boolean;
};

const sourceIdRegex = /^organizations\/[0-9]+\/sources\/[0-9]+$/;

export const validationSchema = yup.object().shape({
    notifier: yup.object().shape({
        name: yup.string().trim().required('Required'),
        cscc: yup.object().shape({
            serviceAccount: yup
                .string()
                .trim()
                .required('A service account is required')
                .test('isValidJson', 'Service account must be valid JSON', (value) => {
                    if (!value) {
                        return false;
                    }
                    try {
                        JSON.parse(value);
                    } catch (e) {
                        return false;
                    }
                    return true;
                }),
            sourceId: yup
                .string()
                .trim()
                .required('A source ID is required')
                .matches(
                    sourceIdRegex,
                    'SCC source ID must match the format: organizations/[0-9]+/sources/[0-9]+'
                ),
        }),
    }),
    updatePassword: yup.bool(),
});

export const defaultValues: GoogleCloudSccIntegrationFormValues = {
    notifier: {
        name: '',
        cscc: {
            serviceAccount: '',
            sourceId: '',
        },
        uiEndpoint: window.location.origin,
        type: 'cscc',
        enabled: true,
    },
    updatePassword: true,
};

function GoogleCloudSccIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<GoogleCloudSccIntegration>): ReactElement {
    const formInitialValues = defaultValues;
    if (initialValues) {
        formInitialValues.notifier = {
            ...formInitialValues.notifier,
            ...initialValues,
        };
        // We want to clear the password because backend returns '******' to represent that there
        // are currently stored credentials
        formInitialValues.notifier.cscc.serviceAccount = '';
    }
    const {
        values,
        touched,
        errors,
        dirty,
        isValid,
        setFieldValue,
        handleBlur,
        isSubmitting,
        isTesting,
        onSave,
        onTest,
        onCancel,
        message,
    } = useIntegrationForm<GoogleCloudSccIntegrationFormValues, typeof validationSchema>({
        initialValues: formInitialValues,
        validationSchema,
    });
    const { isCreating } = usePageState();

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    return (
        <>
            <PageSection variant="light" isFilled hasOverflowScroll>
                {message && <FormMessage message={message} />}
                <Form isWidthLimited>
                    <FormLabelGroup
                        label="Integration name"
                        isRequired
                        fieldId="notifier.name"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="notifier.name"
                            value={values.notifier.name}
                            placeholder="(example, Cloud SCC Integration)"
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    <FormLabelGroup
                        label="Cloud SCC Source ID"
                        isRequired
                        fieldId="notifier.cscc.sourceId"
                        touched={touched}
                        errors={errors}
                    >
                        <TextInput
                            isRequired
                            type="text"
                            id="notifier.cscc.sourceId"
                            value={values.notifier.cscc.sourceId}
                            placeholder="example, organizations/123/sources/456"
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable}
                        />
                    </FormLabelGroup>
                    {!isCreating && (
                        <FormLabelGroup
                            label=""
                            fieldId="updatePassword"
                            helperText="Leave this off to use the currently stored credentials."
                            errors={errors}
                        >
                            <Checkbox
                                label="Update token"
                                id="updatePassword"
                                isChecked={values.updatePassword}
                                onChange={onChange}
                                onBlur={handleBlur}
                                isDisabled={!isEditable}
                            />
                        </FormLabelGroup>
                    )}
                    <FormLabelGroup
                        label="Service Account Key (JSON)"
                        isRequired={values.updatePassword}
                        fieldId="notifier.cscc.serviceAccount"
                        touched={touched}
                        errors={errors}
                    >
                        <TextArea
                            className="json-input"
                            isRequired={values.updatePassword}
                            type="text"
                            id="notifier.cscc.serviceAccount"
                            value={values.notifier.cscc.serviceAccount}
                            placeholder={
                                values.updatePassword
                                    ? 'example,\n{\n  "type": "service_account",\n  "project_id": "123456"\n  ...\n}'
                                    : 'Currently-stored credentials will be used.'
                            }
                            onChange={onChange}
                            onBlur={handleBlur}
                            isDisabled={!isEditable || !values.updatePassword}
                        />
                    </FormLabelGroup>
                </Form>
            </PageSection>
            {isEditable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!dirty || !isValid}
                    >
                        Save
                    </FormSaveButton>
                    <FormTestButton
                        onTest={onTest}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isValid={isValid}
                    >
                        Test
                    </FormTestButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default GoogleCloudSccIntegrationForm;
